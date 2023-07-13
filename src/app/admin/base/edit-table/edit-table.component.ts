import {Component, OnInit, Input, forwardRef} from '@angular/core';
import {FormBuilder, FormArray, ControlValueAccessor, NG_VALUE_ACCESSOR, FormGroup} from "@angular/forms";
import {NzMessageService} from "ng-zorro-antd/message";
import {CdkDragDrop, moveItemInArray} from '@angular/cdk/drag-drop';
import {NzSelectOptionInterface} from "ng-zorro-antd/select";


export interface EditTableItem {
    name: string
    label: string
    type?: string
    placeholder?: string
    default?: any
    options?: NzSelectOptionInterface[]
}


@Component({
    selector: 'app-edit-table',
    templateUrl: './edit-table.component.html',
    styleUrls: ['./edit-table.component.scss'],
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => EditTableComponent),
            multi: true
        }
    ]
})
export class EditTableComponent implements OnInit, ControlValueAccessor {
    group!: FormGroup
    formArray!: FormArray;

    row: any = {};
    _items: any = [];

    onChanged: any = () => {
    }
    onTouched: any = () => {
    }

    @Input()
    set items(data: EditTableItem[]) {
        //TODO 创建默认group
        const row: any = {};
        data.forEach(item => {
            //if (item.hasOwnProperty("default"))
            row[item.name] = item.default
            //TODO 需要补充缺失的控件

        })
        this.row = row;
        this._items = data;
    };

    constructor(
        private msg: NzMessageService,
        private fb: FormBuilder
    ) {
    }

    ngOnInit(): void {
        this.formArray = this.fb.array([]);
        //this.formArray.value
        this.group = this.fb.group({
            array: this.formArray
        })
    }

    writeValue(data: any): void {
        const itemObj = JSON.parse(JSON.stringify(this.row));
        if (data && data.length) {
            data.forEach((item: any) => {
                const newGroup = this.fb.group(Object.assign(itemObj, item));
                this.formArray.push(newGroup);
            });
        }
    }

    registerOnChange(fn: any): void {
        this.onChanged = fn;
    }

    registerOnTouched(fn: any): void {
        this.onTouched = fn;
    }

    change() {
        const data = this.formArray.controls.map((item) => item.value);
        this.onChanged(data);
    }

    handleCopyProperTy(index: number) {
        const old = this.formArray.controls[index].value;
        this.formArray.insert(index, this.fb.group(old));
        this.msg.success("复制成功");
        this.change();
    }

    propertyDel(i: number) {
        this.formArray.removeAt(i);
        this.change();
    }

    drop(event: CdkDragDrop<string[]>): void {
        moveItemInArray(this.formArray.controls, event.previousIndex, event.currentIndex);
        this.change();
    }

    propertyAdd() {
        this.formArray.insert(0, this.fb.group(this.row));
    }
}
