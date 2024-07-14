import {Component, Input, ViewChild} from '@angular/core';
import {SmartEditorComponent, SmartField, SmartRequestService} from "@god-jason/smart";
import {NzNotificationService} from "ng-zorro-antd/notification";
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzCardComponent} from "ng-zorro-antd/card";

@Component({
    selector: 'app-product-action-edit',
    standalone: true,
    imports: [
        SmartEditorComponent,
        NzButtonComponent,
        NzCardComponent,
    ],
    templateUrl: './product-action-edit.component.html',
    styleUrl: './product-action-edit.component.scss'
})
export class ProductActionEditComponent {
    @Input() product_id!: any;

    @ViewChild("editor") editor!: SmartEditorComponent;

    arguments: SmartField[] = [
        {label: '变量', key: 'name', type: 'text'},
        {label: '名称', key: 'label', type: 'text'},
        {
            label: '类型', key: 'type', type: 'select', default: 'int',
            options: [
                {label: '整数', value: 'int'},
                {label: '浮点数', value: 'float'},
                {label: '布尔型', value: 'bool'},
                {label: '字符串', value: 'text'},
                //{label: '枚举', value: 'enum'},
                {label: '数组', value: 'array'},
                {label: '对象', value: 'object'}
            ]
        }
    ]

    data: any = {}
    fields: SmartField[] = [
        {
            label: '', key: 'actions', type: 'list',
            children: [
                {label: '动作', key: 'name', type: 'text'},
                {label: '名称', key: 'label', type: 'text'},
                {label: '异步', key: 'async', type: 'switch'},
                {label: '输入', key: 'inputs', type: 'table', children: this.arguments},
                {label: '输出', key: 'outputs', type: 'table', children: this.arguments},
            ]
        },
    ]


    constructor(private rs: SmartRequestService, private ms: NzNotificationService) {

    }

    ngOnInit(): void {
        this.rs.get(`product/${this.product_id}/config/action`).subscribe(res => {
            this.data = {actions: res.data || []}
        })
    }

    onSubmit() {
        let value = this.editor.value
        this.rs.post(`product/${this.product_id}/config/action`, value.actions).subscribe(res => {
            this.ms.success("提示", "保存成功")
        })
    }

}
