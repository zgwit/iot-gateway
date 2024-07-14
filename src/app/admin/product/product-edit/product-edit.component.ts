import {AfterViewInit, Component, OnInit, TemplateRef, ViewChild} from '@angular/core';
import {NzButtonComponent} from 'ng-zorro-antd/button';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {SmartEditorComponent, SmartField, SmartRequestService} from '@god-jason/smart';
import {NzMessageService} from 'ng-zorro-antd/message';
import {CommonModule} from '@angular/common';
import {NzCardComponent} from "ng-zorro-antd/card";
import {InputProtocolComponent} from "../../../components/input-protocol/input-protocol.component";
import {ReactiveFormsModule} from "@angular/forms";

@Component({
    selector: 'app-product-edit',
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
    templateUrl: './product-edit.component.html',
    styleUrl: './product-edit.component.scss',
})
export class ProductEditComponent implements OnInit, AfterViewInit {
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
            {key: "name", label: "名称", type: "text", required: true, default: '新产品'},
            {key: "keywords", label: "关键字", type: "tags", default: []},
            {key: "icon", label: "图标", type: "text"},
            {key: "protocol", label: "协议", type: "template", template: this.chooseProtocol},
            {key: "url", label: "链接", type: "text"},
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
        this.rs.get(`product/` + this.id).subscribe((res) => {
            this.data = res.data
        });
    }

    onSubmit() {
        if (!this.form.valid) {
            this.msg.error('请检查数据')
            return
        }

        let url = `product/${this.id || 'create'}`
        this.rs.post(url, this.form.value).subscribe((res) => {
            this.router.navigateByUrl('/admin/product/' + res.data.id);
            this.msg.success('保存成功');
        });
    }
}
