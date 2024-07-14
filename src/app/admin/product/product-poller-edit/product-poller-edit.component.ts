import {Component, Input, ViewChild} from '@angular/core';
import {SmartEditorComponent, SmartField, SmartRequestService} from "@god-jason/smart";
import {NzNotificationService} from "ng-zorro-antd/notification";
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzCardComponent} from "ng-zorro-antd/card";

@Component({
    selector: 'app-product-poller-edit',
    standalone: true,
    imports: [
        NzButtonComponent,
        NzCardComponent,
        SmartEditorComponent
    ],
    templateUrl: './product-poller-edit.component.html',
    styleUrl: './product-poller-edit.component.scss'
})
export class ProductPollerEditComponent {
    @Input() product_id!: any;

    @ViewChild("editor") editor!: SmartEditorComponent;

    data: any = {}
    fields: SmartField[] = [
        {label: '', key: 'pollers', type: 'table', children: []},
    ]


    constructor(private rs: SmartRequestService, private ms: NzNotificationService) {

    }

    ngOnInit(): void {
        this.rs.get(`product/${this.product_id}`).subscribe(res => {
            let protocol = res.data.protocol
            this.rs.get(`protocol/${protocol}/poller`).subscribe(res => {
                this.fields[0].children = res.data
                this.rs.get(`product/${this.product_id}/config/poller`).subscribe(res => {
                    this.data = {pollers: res.data || []}
                })
            })
        })
    }

    onSubmit() {
        let value = this.editor.value
        this.rs.post(`product/${this.product_id}/config/poller`, value.pollers).subscribe(res => {
            this.ms.success("提示", "保存成功")
        })
    }


}
