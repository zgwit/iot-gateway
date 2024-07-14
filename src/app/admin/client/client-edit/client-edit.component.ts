import {AfterViewInit, Component, OnInit, TemplateRef, ViewChild} from '@angular/core';
import {NzButtonComponent} from 'ng-zorro-antd/button';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {NzMessageService} from 'ng-zorro-antd/message';
import {CommonModule} from '@angular/common';
import {NzCardComponent} from "ng-zorro-antd/card";
import {InputProtocolComponent} from "../../../components/input-protocol/input-protocol.component";
import {ReactiveFormsModule} from "@angular/forms";
import {SmartEditorComponent, SmartField, SmartRequestService} from "@god-jason/smart";

@Component({
    selector: 'app-clients-edit',
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
    templateUrl: './client-edit.component.html',
    styleUrls: ['./client-edit.component.scss'],
})
export class ClientEditComponent implements OnInit, AfterViewInit {
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
            {key: "name", label: "名称", type: "text", required: true, default: '新客户端'},
            {
                key: "net", label: "网络", type: "select", default: 'tcp', options: [
                    {label: 'TCP', value: 'tcp'},
                    {label: 'UDP', value: 'udp'},
                ]
            },
            {key: "addr", label: "地址", type: "text"},
            {key: "port", label: "端口", type: "number", min: 1, max: 65535},
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
        this.rs.get(`client/` + this.id).subscribe((res) => {
            this.data = res.data
            this.loadProtocolOptions(this.data.protocol_name)
        });
    }

    loadProtocolOptions(protocol: string) {
        //let protocol = this.data.protocol_name || this.form.value.protocol_name
        //this.data = this.form.value //备份数据
        //Object.assign(this.data, this.form.value)
        if (protocol)
            this.rs.get(`protocol/${protocol}/option`).subscribe((res) => {
                this.fields[6].children = res.data
                this.form.group.setControl("protocol_options", this.form.build(res.data, this.form.value.protocol_options))
            });
    }

    onSubmit() {
        if (!this.form.valid) {
            this.msg.error('请检查数据')
            return
        }

        let url = `client/${this.id || 'create'}`
        this.rs.post(url, this.form.value).subscribe((res) => {
            this.router.navigateByUrl('/admin/client/' + res.data.id);
            this.msg.success('保存成功');
        });
    }
}
