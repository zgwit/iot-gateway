import {Component, Input} from '@angular/core';
import {CommonModule} from "@angular/common";
import {NzColDirective, NzRowDirective} from "ng-zorro-antd/grid";
import {NzStatisticComponent} from "ng-zorro-antd/statistic";
import {ActivatedRoute, Router, RouterLink} from "@angular/router";
import {NzMessageService} from "ng-zorro-antd/message";
import {NgxEchartsDirective} from "ngx-echarts";
import {FormBuilder, FormGroup, FormsModule, ReactiveFormsModule} from "@angular/forms";
import {NzFormModule} from "ng-zorro-antd/form";
import {NzOptionComponent, NzSelectComponent} from "ng-zorro-antd/select";
import {NzInputNumberComponent, NzInputNumberGroupComponent} from "ng-zorro-antd/input-number";
import {NzDatePickerModule} from "ng-zorro-antd/date-picker";
import {NzSwitchComponent} from "ng-zorro-antd/switch";
import {SmartRequestService} from "@god-jason/smart";
import * as dayjs from "dayjs";

@Component({
    selector: 'app-device-property',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        NzFormModule,
        NzRowDirective,
        NzColDirective,
        NzStatisticComponent,
        RouterLink,
        NgxEchartsDirective,
        NzSelectComponent,
        NzOptionComponent,
        NzInputNumberGroupComponent,
        NzInputNumberComponent,
        NzDatePickerModule,
        NzSwitchComponent,
        FormsModule,
    ],
    templateUrl: './device-property.component.html',
    styleUrl: './device-property.component.scss'
})
export class DevicePropertyComponent {
    group!: FormGroup;

    data: any = {};

    @Input() id!: any;
    properties: any = [];

    values: any = {};

    actives: any = {}
    names: any = {}

    loading = false;


    query: any = {}
    chart: any;

    option: any = {
        tooltip: {
            trigger: 'axis',
            position: function (pt: any) {
                return [pt[0], '10%'];
            },
        },
        title: {left: 'center', text: '历史曲线',},
        toolbox: {feature: {saveAsImage: {}},},
        xAxis: {type: 'time', boundaryGap: false},
        yAxis: {type: 'value'},
        dataZoom: [
            {type: 'inside', start: 0, end: 100,},
            {start: 0, end: 100,},
        ],
        series: [],
    };

    constructor(
        private router: Router,
        private msg: NzMessageService,
        private fb: FormBuilder,
        private rs: SmartRequestService,
        private route: ActivatedRoute
    ) {
        //this.load();
    }

    chartInit(ec: any) {
        this.chart = ec
    }

    ngOnInit(): void {
        this.load();
        this.loadValues();

        this.group = this.fb.group({
            strEnd: [[dayjs().add(-7, "days").toDate(), dayjs().toDate(),],],
            window: [1],
            winTp: ['h'],
            fn: ['last'],
        });
    }

    load() {
        this.rs.get(`device/${this.id}`, {}).subscribe((res) => {
            this.data = res.data;
            this.loadProperties();
        });
    }

    loadValues() {
        this.rs.get('device/' + this.id + '/values').subscribe((res) => {
            this.values = res.data || {};
        });
    }

    loadProperties() {
        this.rs.get(`product/${this.data.product_id}/config/property`).subscribe(res => {
            this.properties = res.data
            this.properties.forEach((p: any) => {
                this.names[p.name] = p.label || p.name
            })
        })
    }

    loadHistory(name: string, query: any) {
        this.rs.get(`device/${this.id}/history/${name}`, query).subscribe((res) => {
            console.log("history", res.data)
            //this.searchData = res.data || [];
            //this.searchTotal = res.total || 0;
            //图表渲染
            // this.chart.setOption(this.option);
            this.option.series.push({
                name: this.names[name], //使用名称
                type: 'line',
                smooth: true,
                symbol: 'none',
                //areaStyle: {},
                data: res.data?.map((p: any) => [p.time, p.value]),
            })
            // this.option.xAxis.data = res.data?.map((p: any) => dayjs(p.time).format())
            this.chart.clear()
            this.chart.setOption(this.option)
            console.log(this.option)
        })
    }


    search() {
        if (this.group.valid) {
            let value = this.group.value;
            let query = {
                start: dayjs(value.strEnd[0]).toISOString(),
                end: dayjs(value.strEnd[1]).toISOString(),
                window: value.window ? value.window + value.winTp : '1h',
                fn: 'last',
            };

            console.log('actives', this.actives)

            this.option.series = []
            Object.keys(this.actives).forEach((k: string) => {
                if (this.actives[k])
                    this.loadHistory(k, query)
            })
            //this.loadHistory("a", query)
        } else {
            Object.values(this.group.controls).forEach((control) => {
                if (control.invalid) {
                    control.markAsDirty();
                    control.updateValueAndValidity({onlySelf: true});
                }
            });
        }
    }

    valueSwitch(name: string, $event: any) {
        //console.log(name, $event)
        this.rs.post('device/' + this.id + '/data', {[name]: $event}).subscribe((res) => {
            this.msg.success("操作成功")
        });
    }
}
