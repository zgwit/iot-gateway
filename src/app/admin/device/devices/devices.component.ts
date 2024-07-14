import {Component, Inject, Input, OnInit, Optional,} from '@angular/core';
import {ActivatedRoute,} from '@angular/router';
import {CommonModule} from '@angular/common';
import {
    ParamSearch,
    SmartRequestService,
    SmartTableButton,
    SmartTableColumn,
    SmartTableComponent,
    SmartTableOperator
} from '@god-jason/smart';
import {NZ_MODAL_DATA, NzModalRef} from 'ng-zorro-antd/modal';

@Component({
    selector: 'app-devices',
    standalone: true,
    imports: [CommonModule, SmartTableComponent],
    templateUrl: './devices.component.html',
    styleUrl: './devices.component.scss',
})
export class DevicesComponent implements OnInit {
    @Input() product_id: any = '';
    @Input() tunnel_id: any = '';


    datum: any[] = [];
    total = 0;
    loading = false;


    buttons: SmartTableButton[] = [
        {icon: 'plus', label: '创建', link: () => `admin/device/create`},
    ];

    columns: SmartTableColumn[] = [
        {
            key: 'id', sortable: true, label: 'ID', keyword: true,
            link: (data) => `admin/device/${data.id}`,
        },
        {key: 'name', sortable: true, label: '名称', keyword: true},
        {
            key: 'product', sortable: true, label: '产品', keyword: true,
            link: (data) => `admin/product/${data.product_id}`,
        },
        {key: 'online', sortable: true, label: '上线时间', date: true},
    ];

    columnsSelect: SmartTableColumn[] = [
        {key: 'id', label: 'ID', keyword: true},
        {key: 'name', label: '名称', keyword: true},
    ];

    operators: SmartTableOperator[] = [
        {icon: 'edit', title: '编辑', link: (data) => `admin/device/${data.id}/edit`,},
        {
            icon: 'delete', title: '删除', confirm: '确认删除？', action: (data) => {
                this.rs.get(`device/${data.id}/delete`).subscribe((res) => this.refresh())
            },
        },
    ];

    operatorsSelect: SmartTableOperator[] = [
        {label: '选择', action: (data) => this.ref.close(data)},
    ];
    query!: ParamSearch

    constructor(
        private route: ActivatedRoute,
        private rs: SmartRequestService,
        @Optional() protected ref: NzModalRef,
        @Optional() @Inject(NZ_MODAL_DATA) protected data: any
    ) {
    }

    ngOnInit(): void {
    }

    refresh() {
        this.search(this.query)
    }

    search(query: ParamSearch) {
        //console.log('onQuery', query)        this.query = query
        if (this.product_id) query.filter['product_id'] = this.product_id;
        if (this.tunnel_id) query.filter['tunnel_id'] = this.tunnel_id;

        this.loading = true;
        this.rs.post('device/search', query).subscribe((res) => {
            this.datum = res.data;
            this.total = res.total;
        }).add(() => this.loading = false);
    }
}
