import {Component, Input, Optional} from '@angular/core';
import {
    ParamSearch,
    SmartRequestService,
    SmartTableButton,
    SmartTableColumn,
    SmartTableComponent,
    SmartTableOperator
} from '@god-jason/smart';
import {NzModalRef} from 'ng-zorro-antd/modal';
import {CommonModule} from '@angular/common';

@Component({
    selector: 'app-links',
    standalone: true,
    imports: [
        CommonModule,
        SmartTableComponent,
    ],
    templateUrl: './links.component.html',
    styleUrls: ['./links.component.scss'],
})
export class LinksComponent {
    @Input() server_id!: string;

    datum: any[] = [];
    total = 0;
    loading = false;

    buttons: SmartTableButton[] = [
        {icon: "plus", label: "创建", link: () => `/admin/link/create`}
    ];

    columns: SmartTableColumn[] = [
        {key: "id", sortable: true, label: "ID", keyword: true, link: (data) => `/admin/link/${data.id}`},
        {
            key: "server",
            sortable: true,
            label: "服务器",
            keyword: true,
            link: (data) => `/admin/server/${data.server_id}`
        },
        {key: "name", sortable: true, label: "名称", keyword: true},
        {key: "remote", sortable: true, label: "远程地址", keyword: true},
        {key: "status", label: "状态"},
        {key: "created", sortable: true, label: "创建时间", date: true},
    ];

    columnsSelect: SmartTableColumn[] = [
        {key: "id", label: "ID", keyword: true},
        {key: "server", label: "服务器", keyword: true},
        {key: "name", label: "名称", keyword: true},
        {key: "remote", label: "远程地址", keyword: true},
        {key: "status", label: "状态"},
    ];

    operators: SmartTableOperator[] = [
        {icon: 'edit', title: '编辑', link: data => `/admin/link/${data.id}/edit`},
        {
            icon: 'delete', title: '删除', confirm: "确认删除？", action: data => {
                this.rs.get(`link/${data.id}/delete`).subscribe(res => this.refresh())
            }
        },
        {
            icon: 'close-circle', title: '停止', action: data => {
                this.rs.get(`link/${data.id}/close`).subscribe(res => this.refresh())
            }
        }
    ];

    operatorsSelect: SmartTableOperator[] = [
        {label: "选择", action: data => this.ref.close(data)},
    ];
    query!: ParamSearch

    constructor(private rs: SmartRequestService, @Optional() protected ref: NzModalRef) {
    }

    refresh() {
        this.search(this.query)
    }

    search(query: ParamSearch) {
        //console.log('onQuery', query)
        this.query = query

        if (this.server_id)
            query.filter['server_id'] = this.server_id;

        this.loading = true
        this.rs.post('link/search', query).subscribe((res) => {
            this.datum = res.data;
            this.total = res.total;
        }).add(() => this.loading = false);
    }

}
