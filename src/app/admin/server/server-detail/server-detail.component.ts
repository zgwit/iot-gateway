import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router, RouterLink} from "@angular/router";
import {NzTabComponent, NzTabDirective, NzTabSetComponent} from "ng-zorro-antd/tabs";
import {CommonModule} from "@angular/common";
import {NzCardComponent} from "ng-zorro-antd/card";
import {NzSpaceComponent, NzSpaceItemDirective} from "ng-zorro-antd/space";
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzPopconfirmDirective} from "ng-zorro-antd/popconfirm";
import {NzMessageService} from "ng-zorro-antd/message";
import {SmartInfoComponent, SmartInfoItem, SmartRequestService} from "@god-jason/smart";
import {LinksComponent} from "../../link/links/links.component";

@Component({
    selector: 'app-servers-detail',
    templateUrl: './server-detail.component.html',
    standalone: true,
    imports: [
        CommonModule,
        NzTabSetComponent,
        NzTabComponent,
        NzTabDirective,
        NzCardComponent,
        SmartInfoComponent,
        NzSpaceComponent,
        NzSpaceItemDirective,
        NzButtonComponent,
        RouterLink,
        NzPopconfirmDirective,
        LinksComponent
    ],
    styleUrls: ['./server-detail.component.scss']
})
export class ServerDetailComponent implements OnInit {
    id: string = ""
    data: any = {}


    fields: SmartInfoItem[] = [
        {key: 'id', label: 'ID'},
        {key: 'name', label: '名称'},
        {key: 'port', label: '端口'},
        {key: 'mode', label: '模式'},
        {key: 'disabled', label: '禁用'},
        {key: 'created', label: '创建时间', type: 'date'},
        {key: "protocol_name", label: "通讯协议"},
        {key: 'description', label: '说明', span: 2},
    ];

    constructor(
        private router: Router,
        private msg: NzMessageService,
        private rs: SmartRequestService,
        private route: ActivatedRoute
    ) {
    }

    ngOnInit(): void {
        // @ts-ignore
        this.id = this.route.snapshot.paramMap.get("id")

        this.load()
    }

    load() {
        this.rs.get(`server/${this.id}`).subscribe(res => {
            this.data = res.data;
        })
    }

    delete() {
        this.rs.get(`server/${this.id}/delete`, {}).subscribe((res: any) => {
            this.msg.success('删除成功');
            this.router.navigateByUrl('/admin/server');
        });
    }
}
