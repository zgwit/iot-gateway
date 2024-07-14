import {AfterViewInit, Component, OnInit, TemplateRef, ViewChild} from '@angular/core';
import {NzButtonComponent} from 'ng-zorro-antd/button';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {NzMessageService} from 'ng-zorro-antd/message';
import {CommonModule} from '@angular/common';
import {NzCardComponent} from "ng-zorro-antd/card";
import {SmartEditorComponent, SmartField, SmartRequestService} from "@god-jason/smart";
import {InputServerComponent} from "../../../components/input-server/input-server.component";
import {InputProtocolComponent} from "../../../components/input-protocol/input-protocol.component";
import {ReactiveFormsModule} from "@angular/forms";

@Component({
    selector: 'app-links-edit',
    standalone: true,
    imports: [
        CommonModule,
        ReactiveFormsModule,
        NzButtonComponent,
        RouterLink,
        NzCardComponent,
        SmartEditorComponent,
        InputServerComponent,
        InputProtocolComponent,
    ],
    templateUrl: './link-edit.component.html',
    styleUrls: ['./link-edit.component.scss'],
})
export class LinkEditComponent implements OnInit, AfterViewInit {
    id: any = '';

    @ViewChild('form') form!: SmartEditorComponent
    @ViewChild("chooseServer") chooseServer!: TemplateRef<any>
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
            {key: "name", label: "名称", type: "text", required: true},
            {key: "server_id", label: "服务器", type: "template", template: this.chooseServer},

            {
                key: "protocol_name", label: "通讯协议", type: "template", template: this.chooseProtocol,
                change: ($event) => setTimeout(() => this.loadProtocolOptions($event))
            },
            {key: "protocol_options", label: "通讯协议参数", type: "object"},
            {key: "description", label: "说明", type: "textarea"},
        ]
    }

    ngAfterViewInit(): void {
        setTimeout(() => this.build(), 1)
    }

    ngOnInit(): void {
        if (this.route.snapshot.paramMap.has('id')) {
            this.id = this.route.snapshot.paramMap.get('id');
            this.load()
        }
    }

    load() {
        this.rs.get(`link/` + this.id).subscribe((res) => {
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

        let url = `link/${this.id || 'create'}`
        this.rs.post(url, this.form.value).subscribe((res) => {
            this.router.navigateByUrl('/admin/link/' + res.data.id);
            this.msg.success('保存成功');
        });
    }
}
