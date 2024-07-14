import {Component, OnInit} from '@angular/core';
import {NzSpaceComponent, NzSpaceItemDirective} from 'ng-zorro-antd/space';
import {NzButtonComponent} from 'ng-zorro-antd/button';
import {CommonModule} from '@angular/common';
import {ActivatedRoute, Router, RouterLink} from '@angular/router';
import {NzPopconfirmDirective} from 'ng-zorro-antd/popconfirm';
import {SmartInfoComponent, SmartInfoItem, SmartRequestService} from '@god-jason/smart';
import {NzMessageService} from 'ng-zorro-antd/message';
import {NzCardComponent} from "ng-zorro-antd/card";
import {NzTabsModule} from "ng-zorro-antd/tabs";
import {DevicesComponent} from "../../device/devices/devices.component";
import {NzCollapseComponent} from "ng-zorro-antd/collapse";
import {ProductActionEditComponent} from "../product-action-edit/product-action-edit.component";
import {ProductMapperEditComponent} from "../product-mapper-edit/product-mapper-edit.component";
import {ProductPollerEditComponent} from "../product-poller-edit/product-poller-edit.component";

@Component({
    selector: 'app-product-detail',
    standalone: true,
    imports: [
        CommonModule,
        SmartInfoComponent,
        NzCardComponent,
        NzSpaceComponent,
        NzButtonComponent,
        NzSpaceItemDirective,
        RouterLink,
        NzPopconfirmDirective,
        NzTabsModule,
        DevicesComponent,
        NzCollapseComponent,
        ProductActionEditComponent,
        ProductMapperEditComponent,
        ProductPollerEditComponent,
    ],
    templateUrl: './product-detail.component.html',
    styleUrl: './product-detail.component.scss'
})
export class ProductDetailComponent implements OnInit {
    id!: any;

    value = '';

    data: any = {};

    fields: SmartInfoItem[] = [
        {label: 'ID', key: 'id'},
        {label: '名称', key: 'name'},
        {label: '关键字', key: 'keywords', type: 'tags'},
        {label: '创建时间', key: 'created', type: 'date'},
        {label: '说明', key: 'description', span: 2},
    ];

    constructor(
        private router: Router,
        private msg: NzMessageService,
        private rs: SmartRequestService,
        private route: ActivatedRoute
    ) {
    }

    ngOnInit(): void {
        if (this.route.snapshot.paramMap.has('id')) {
            this.id = this.route.snapshot.paramMap.get('id');
            this.load()
        }
    }

    load() {
        this.rs.get(`product/${this.id}`, {}).subscribe((res) => {
                this.data = res.data
            }
        );
    }

    delete() {
        this.rs.get(`product/${this.id}/delete`, {}).subscribe((res) => {
            this.msg.success('删除成功');
            this.router.navigateByUrl('admin/product');
        });
        this.load();
    }
}
