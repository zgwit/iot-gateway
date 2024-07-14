import {AfterViewInit, Component, OnInit, TemplateRef, ViewChild} from '@angular/core';
import {NzButtonComponent} from 'ng-zorro-antd/button';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {NzMessageService} from 'ng-zorro-antd/message';
import {CommonModule} from '@angular/common';
import {NzCardComponent} from "ng-zorro-antd/card";
import {SmartEditorComponent, SmartField, SmartRequestService, SmartSelectOption} from "@god-jason/smart";
import {InputProtocolComponent} from "../../../components/input-protocol/input-protocol.component";
import {ReactiveFormsModule} from "@angular/forms";

@Component({
    selector: 'app-serials-edit',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        NzButtonComponent,
        RouterLink,
        NzCardComponent,
        SmartEditorComponent,
        InputProtocolComponent,
    ],
    templateUrl: './serial-edit.component.html',
    styleUrls: ['./serial-edit.component.scss'],
})
export class SerialEditComponent implements OnInit, AfterViewInit {
    id: any = '';

    @ViewChild('form') form!: SmartEditorComponent
    @ViewChild('chooseProtocol') chooseProtocol!: TemplateRef<any>


    ports: SmartSelectOption[] = []

    fields: SmartField[] = []
    data: any = {}

    constructor(private router: Router,
                private msg: NzMessageService,
                private rs: SmartRequestService,
                private route: ActivatedRoute
    ) {
    }

    build() {
        this.fields = [
            {key: "id", label: "ID", type: "text", min: 2, max: 30, placeholder: "选填"},
            {key: "name", label: "名称", type: "text", required: true, default: '新串口'},
            {key: "port_name", label: "端口", type: "select", options: this.ports},
            {
                key: "baud_rate", label: "波特率", type: "select", default: 9600, options: [
                    {label: '150', value: 150},
                    {label: '200', value: 200},
                    {label: '300', value: 300},
                    {label: '600', value: 600},
                    {label: '1200', value: 1200},
                    {label: '1800', value: 1800},
                    {label: '2400', value: 2400},
                    {label: '4800', value: 4800},
                    {label: '9600', value: 9600},
                    {label: '19200', value: 19200},
                    {label: '38400', value: 38400},
                    {label: '57600', value: 57600},
                    {label: '115200', value: 115200},
                ]
            },
            {
                key: "data_bits", label: "字长", type: "radio", options: [
                    {label: '5', value: 5},
                    {label: '6', value: 6},
                    {label: '7', value: 7},
                    {label: '8', value: 8},
                ], default: 8
            },
            {
                key: "parity_mode", label: "校验", type: "radio", options: [
                    {label: '无', value: 0},
                    {label: '奇', value: 1},
                    {label: '偶', value: 2},
                    {label: '1', value: 3},
                    {label: '0', value: 4},
                ]
            },
            {
                key: "stop_bits", label: "停止位", type: "radio", options: [
                    {label: '1', value: 1},
                    {label: '1.5', value: 1.5, disabled: true},
                    {label: '2', value: 2},
                ]
            },
            {
                key: "protocol_name", label: "通讯协议", type: "template", template: this.chooseProtocol,
                change: ($event) => setTimeout(() => this.loadProtocolOptions($event))
            },
            {key: "protocol_options", label: "通讯协议参数", type: "object"},
            {key: "description", label: "说明", type: "textarea"},
        ]
    }

    ngOnInit(): void {
        if (this.route.snapshot.paramMap.has('id')) {
            this.id = this.route.snapshot.paramMap.get('id');
            this.load()
        }
        this.loadPorts()
    }

    ngAfterViewInit(): void {
        setTimeout(() => this.build(), 1)
    }

    load() {
        this.rs.get(`serial/` + this.id).subscribe((res) => {
            this.data = res.data
            this.loadProtocolOptions(this.data.protocol_name)
        });
    }

    loadPorts() {
        this.rs.get(`serial/ports`).subscribe((res) => {
            this.fields[2].options = res.data.map((p: string) => {
                return {value: p, label: p}
            })
        });
    }

    loadProtocolOptions(protocol: string) {
        if (protocol)
            this.rs.get(`protocol/${protocol}/option`).subscribe((res) => {
                this.fields[8].children = res.data
                this.form.group.setControl("protocol_options", this.form.build(res.data, this.form.value.protocol_options))
            });
    }

    onSubmit() {
        if (!this.form.valid) {
            this.msg.error('请检查数据')
            return
        }

        let url = `serial/${this.id || 'create'}`
        this.rs.post(url, this.form.value).subscribe((res) => {
            this.router.navigateByUrl('/admin/serial/' + res.data.id);
            this.msg.success('保存成功');
        });
    }
}
