import {AfterViewInit, Component, OnInit, TemplateRef, ViewChild} from '@angular/core';
import {NzButtonComponent} from 'ng-zorro-antd/button';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {NzMessageService} from 'ng-zorro-antd/message';
import {CommonModule} from '@angular/common';
import {NzCardComponent} from "ng-zorro-antd/card";
import {SmartEditorComponent, SmartField, SmartRequestService} from "@god-jason/smart";
import {InputProtocolComponent} from "../../../components/input-protocol/input-protocol.component";
import {ReactiveFormsModule} from "@angular/forms";

@Component({
    selector: 'app-servers-edit',
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
    templateUrl: './server-edit.component.html',
    styleUrls: ['./server-edit.component.scss'],
})
export class ServerEditComponent implements OnInit, AfterViewInit {
    id: any = '';

    @ViewChild('form') form!: SmartEditorComponent
    @ViewChild('chooseProtocol') chooseProtocol!: TemplateRef<any>


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
            {key: "name", label: "名称", type: "text", required: true, default: '新服务端'},
            {key: "port", label: "端口", type: "number", min: 1, max: 65535, default: 60000},
            {
                key: "mode", label: "模式", type: "select", default: "register", options: [
                    {label: "注册包", value: "register"},
                    {label: "多连接", value: "multiple"},
                    {label: "单连接", value: "single"},
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
    }

    ngAfterViewInit(): void {
        setTimeout(() => this.build(), 1)
    }

    load() {
        this.rs.get(`server/` + this.id).subscribe((res) => {
            this.data = res.data
            this.loadProtocolOptions(this.data.protocol_name)
        });
    }


    loadProtocolOptions(protocol: string) {
        if (protocol)
            this.rs.get(`protocol/${protocol}/option`).subscribe((res) => {
                this.fields[4].children = res.data
                this.form.group.setControl("protocol_options", this.form.build(res.data, this.form.value.protocol_options))
            });
    }

    onSubmit() {
        if (!this.form.valid) {
            this.msg.error('请检查数据')
            return
        }

        let url = `server/${this.id || 'create'}`
        this.rs.post(url, this.form.value).subscribe((res) => {
            this.router.navigateByUrl('/admin/server/' + res.data.id);
            this.msg.success('保存成功');
        });
    }
}
