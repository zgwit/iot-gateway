import {Component} from '@angular/core';
import {NzContextMenuService, NzDropDownModule} from 'ng-zorro-antd/dropdown';
import {NzLayoutModule} from "ng-zorro-antd/layout";
import {RouterModule} from "@angular/router";
import {CommonModule} from "@angular/common";

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
    activeMenu: string = '';
    ary = [false, false
        , false
        , false
        , false]
    menuList = [{
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
    }
        , {
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
        }
        // ,
        // {
        //   title: '系统设置',
        //   icon: 'apartment',
        //   open: false,
        //   children: [
        //     {
        //       title: '网站',
        //       path: '/admin/setting/web'
        //     },
        //     {
        //       title: '数据库',
        //       path: '/admin/setting/database'
        //     },
        //     {
        //       title: '日志',
        //       path: '/admin/setting/log'
        //     },
        //     {
        //       title: '消息总线',
        //       path: '/admin/setting/mqtt'
        //     },
        //   ]
        // }
    ]

    constructor(
        private nzContextMenuService: NzContextMenuService,
    ) {
        for (let index = 0; index < this.menuList.length; index++) {
            const item = this.menuList[index];
            const {children} = item;
            for (let i = 0; i < children.length; i++) {
                const it = children[i];
                if (it.path === location.pathname) {
                    item.open = true;
                }
            }
        }
    }
    // contextMenu($event: MouseEvent, menu: NzDropdownMenuComponent, mes: number): void {
    //   this.edit = mes
    //   this.nzContextMenuService.create($event, menu);
    // }
    // clientFm(num: number) {
    //   this.ary[num] = false

    // }
    selectDropdown(): void {
        this.ary[this.edit] = true
    }
}
