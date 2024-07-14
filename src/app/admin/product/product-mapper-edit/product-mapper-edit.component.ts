import {Component, Input, OnInit, ViewChild} from '@angular/core';
import {SmartEditorComponent, SmartField, SmartRequestService} from "@god-jason/smart";
import {NzNotificationService} from "ng-zorro-antd/notification";
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzCardComponent} from "ng-zorro-antd/card";

@Component({
    selector: 'app-product-mapper-edit',
    standalone: true,
    imports: [
        NzButtonComponent,
        NzCardComponent,
        SmartEditorComponent
    ],
    templateUrl: './product-mapper-edit.component.html',
    styleUrl: './product-mapper-edit.component.scss'
})
export class ProductMapperEditComponent implements OnInit {
    @Input() product_id!: any;

    @ViewChild("editor") editor!: SmartEditorComponent;

    data: any = {}
    fields: SmartField[] = []


    constructor(private rs: SmartRequestService, private ms: NzNotificationService) {

    }

    ngOnInit(): void {
        this.rs.get(`product/${this.product_id}`).subscribe(res => {
            let protocol = res.data.protocol
            this.rs.get(`protocol/${protocol}/mapper`).subscribe(res => {
                this.fields = res.data;
                //this.fields[0].children = res.data
                this.rs.get(`product/${this.product_id}/config/mapper`).subscribe(res => {
                    //this.data = {mappers: res.data || []}
                    this.data = res.data || {}
                })
            })
        })
    }

    onSubmit() {
        let value = this.editor.value
        this.rs.post(`product/${this.product_id}/config/mapper`, value.mappers).subscribe(res => {
            this.ms.success("提示", "保存成功")
        })
    }

}
