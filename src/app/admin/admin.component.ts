import {Component} from '@angular/core';
import {NzDropDownModule} from 'ng-zorro-antd/dropdown';
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {RouterModule} from "@angular/router";
import {CommonModule} from "@angular/common";
import {SmartRequestService} from "@god-jason/smart";

@Component({
    selector: 'app-admin',
    templateUrl: './admin.component.html',
    standalone: true,
    imports: [
        CommonModule,
        NzLayoutModule,
        NzDropDownModule,
        RouterModule,
    ],
    styleUrls: ['./admin.component.scss']
})
export class AdminComponent {
    edit!: number;

    menuSetting: any = {
        title: '系统设置',
        icon: 'apartment',
        open: false,
        children: []
    }

    menuList: any = [
        {
            title: '产品管理',
            icon: 'project',
            open: false,
            children: [
                {
                    title: '所有产品',
                    path: '/admin/product'
                },
                {
                    title: '所有设备',
                    path: '/admin/device'
                }
            ]
        },
        {
            title: '连接管理',
            icon: 'appstore',
            open: false,
            children: [
                {
                    title: 'TCP客户端',
                    path: '/admin/client'
                },
                {
                    title: 'TCP服务器',
                    path: '/admin/server'
                },
                {
                    title: 'TCP连接',
                    path: '/admin/link'
                },
                {
                    title: '串口连接',
                    path: '/admin/serial'
                },
            ]
        },
        this.menuSetting
    ]

    constructor(private rs: SmartRequestService,) {
        this.load()
    }

    load() {
        this.rs.get('setting/modules').subscribe(res => {
            res.data?.forEach((s: any) => {
                this.menuSetting.children.push({
                    name: s.name,
                    path: "/admin/setting",
                    query: {module: s.module}
                })
            })
        })
    }
}
