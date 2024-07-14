import {Component, Optional} from '@angular/core';
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
    selector: 'app-servers',
    standalone: true,
    imports: [
        CommonModule,
        SmartTableComponent,
    ],
    templateUrl: './servers.component.html',
    styleUrls: ['./servers.component.scss'],
})
export class ServersComponent {
    datum: any[] = [];
    total = 0;
    loading = false;

    buttons: SmartTableButton[] = [
        {icon: "plus", label: "创建", link: () => `/admin/server/create`}
    ];

    columns: SmartTableColumn[] = [
        {key: "id", sortable: true, label: "ID", keyword: true, link: (data) => `/admin/server/${data.id}`},
        {key: "name", sortable: true, label: "名称", keyword: true},
        {key: "port", sortable: true, label: "端口", keyword: true},
        {key: "mode", sortable: true, label: "模式", keyword: true},
        {key: "status", label: "状态"},
        {key: "created", sortable: true, label: "创建时间", date: true},
    ];

    columnsSelect: SmartTableColumn[] = [
        {key: "id", label: "ID", keyword: true},
        {key: "name", label: "名称", keyword: true},
        {key: "port", label: "端口", keyword: true},
        {key: "mode", label: "模式", keyword: true},
        {key: "status", label: "状态"},
    ];

    operators: SmartTableOperator[] = [
        {icon: 'edit', title: '编辑', link: data => `/admin/server/${data.id}/edit`},
        {
            icon: 'delete', title: '删除', confirm: "确认删除？", action: data => {
                this.rs.get(`server/${data.id}/delete`).subscribe(res => this.refresh())
            }
        },
        {
            icon: 'play-circle', title: '启动', action: data => {
                this.rs.get(`server/${data.id}/open`).subscribe(res => this.refresh())
            }
        },
        {
            icon: 'close-circle', title: '停止', action: data => {
                this.rs.get(`server/${data.id}/close`).subscribe(res => this.refresh())
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
        this.loading = true
        this.rs.post('server/search', query).subscribe((res) => {
            this.datum = res.data;
            this.total = res.total;
        }).add(() => this.loading = false);
    }

}
