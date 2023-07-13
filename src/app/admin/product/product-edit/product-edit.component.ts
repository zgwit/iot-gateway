import {RequestService} from '../../../request.service';
import {Component, Input, Output, EventEmitter, OnInit} from '@angular/core';
import {
    UntypedFormBuilder,
    UntypedFormControl,
    FormGroup,
    UntypedFormGroup,
    ValidationErrors,
    Validators,
    FormsModule,
    FormBuilder,
} from '@angular/forms';
import {NzMessageService} from 'ng-zorro-antd/message';
import {DatePipe} from '@angular/common';
import {CdkDragDrop, moveItemInArray} from '@angular/cdk/drag-drop';
import {ActivatedRoute, Router} from '@angular/router';
import {EditTableItem} from "../../base/edit-table/edit-table.component";

@Component({
    selector: 'app-product-edit',
    templateUrl: './product-edit.component.html',
    styleUrls: ['./product-edit.component.scss'],
    providers: [DatePipe],
})
export class ProductEditComponent implements OnInit {
    validateForm!: any;
    id: any = 0;

    listData: EditTableItem[] = [{
        label: '名称',
        name: 'name'
    }, {
        label: '类型',
        name: 'type',
        type: 'select',
        default: 'uint16',
        options: [{
            label: '位 BIT',
            value: 'bit'
        }, {
            label: '短整数 INT16',
            value: 'int16'
        }, {
            label: '无符号短整数 UINT16',
            value: 'uint16'
        }, {
            label: '整数 INT32',
            value: 'int32'
        }, {
            label: '无符号整数 UINT32',
            value: 'uint32'
        }, {
            label: '浮点数 FLOAT',
            value: 'float'
        }, {
            label: '双精度浮点数 DOUBLE',
            value: 'double'
        }]
    }, {
        label: '偏移',
        name: 'offset',
        type: 'number',
        default: 0
    }, {
        label: '位',
        name: 'bits',
        type: 'number',
        default: 0
    }, {
        label: '大端',
        name: 'be',
        type: 'switch',
        default: true
    }, {
        label: '倍率',
        name: 'rate',
        type: 'number',
        default: 1
    }]


    listDataBits: EditTableItem[] = [{
        label: '名称',
        name: 'name'
    }, {
        label: '偏移',
        name: 'offset',
        type: 'number',
        default: 0
    }]


    listFilter = [{
        label: '过滤字段',
        name: 'name',
        placeholder: '例 a',
    }, {
        label: '表达式（条件为 假 时过滤掉）',
        name: 'expression',
        placeholder: '例 a<1 && a<10',
    }]

    listCalculators = [{
        label: '赋值字段',
        name: 'name',
        placeholder: '例 c',
    }, {
        label: '表达式',
        name: 'expression',
        placeholder: '例 a+b',
    }]


    mode = "new";

    constructor(
        private readonly datePipe: DatePipe,
        private fb: FormBuilder,
        private msg: NzMessageService,
        private rs: RequestService,
        private route: ActivatedRoute,
        private router: Router
    ) {
        this.build();
    }

    ngOnInit(): void {
        if (this.route.snapshot.paramMap.has('id')) {
            this.mode = "edit";
            this.id = this.route.snapshot.paramMap.get('id');
            this.rs.get(`product/${this.id}`).subscribe((res) => {
                this.build(res.data);
            });
        }
        this.build();
    }

    build(obj?: any) {
        obj = obj || {};
        this.validateForm = this.fb.group({
            id: [obj.id || '', []],
            name: [obj.name || '', [Validators.required]],
            desc: [obj.desc || '', []],
            mappers: this.fb.array(
                obj.mappers
                    ? obj.mappers.map((prop: any) =>
                        this.fb.group({
                            code: [prop.code || 3, []],
                            addr: [prop.addr || 0, []],
                            size: [prop.size || 0, []],
                            points: [prop.points || [], []],
                        })) : []),
            filters: [obj.filters || [], []],
            calculators: [obj.calculators || [], []],
        })
    }


    handleCancel() {
        this.router.navigateByUrl(`/admin/product`);
    }

    submit() {
        this.validateForm.updateValueAndValidity();
        if (this.validateForm.valid) {
            let url = this.id ? `product/${this.id}` : `product/create`;
            if (this.mode === "edit" && !this.validateForm.value.id) {
                this.msg.warning('ID不可为空');
                return;
            }
            this.rs.post(url, this.validateForm.value).subscribe((res) => {
                this.msg.success('保存成功');
                this.router.navigateByUrl(`/admin/product`);
            });
            return;
        } else {
            Object.values(this.validateForm.controls).forEach((control: any) => {
                if (control.invalid) {
                    control.markAsDirty();
                    control.updateValueAndValidity({onlySelf: true});
                }
            });
        }
    }

    reset() {
        this.validateForm.reset();
        for (const key in this.validateForm.controls) {
            if (this.validateForm.controls.hasOwnProperty(key)) {
                this.validateForm.controls[key].markAsPristine();
                this.validateForm.controls[key].updateValueAndValidity();
            }
        }
    }

    addMapper() {
        this.validateForm.get('mappers').push(
            this.fb.group({
                code: [3, []],
                addr: [0, []],
                size: [0, []],
                points: [[], []],
            })
        );
    }

    drop(mapper
             :
             any, event
             :
             CdkDragDrop<string[]>
    ):
        void {
        moveItemInArray(
            mapper.get('points').controls,
            event.previousIndex,
            event.currentIndex
        )
        ;
    }

    pointCopy(mapper
                  :
                  any, index
                  :
                  number
    ) {
        const oitem = mapper.get('points').controls[index].value;
        mapper.get('points').insert(index, this.fb.group(oitem));
        this.msg.success('复制成功');
    }

    pointDel(mapper
                 :
                 any, i
                 :
                 number
    ) {
        mapper.get('points').removeAt(i);
    }

    mapperDel(i
                  :
                  number
    ) {
        this.validateForm.get('mappers').removeAt(i);
    }

    pointAdd(mapper
                 :
                 any
    ) {
        mapper.get('points').push(
            this.fb.group({
                name: ['', []],
                type: ['word', []],
                offset: [0, []],
                be: [true, []],
                rate: [1, []],
            })
        );
    }
}
