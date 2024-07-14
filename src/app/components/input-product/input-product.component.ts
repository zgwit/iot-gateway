import {Component, forwardRef, Input, OnInit} from '@angular/core';
import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/forms";
import {NzInputDirective} from "ng-zorro-antd/input";
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzModalService} from "ng-zorro-antd/modal";
import {SmartRequestService} from "@god-jason/smart";
import {ProductsComponent} from "../../admin/product/products/products.component";

@Component({
    selector: 'app-input-product',
    standalone: true,
    imports: [
        NzInputDirective,
        NzButtonComponent
    ],
    templateUrl: './input-product.component.html',
    styleUrl: './input-product.component.scss',
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => InputProductComponent),
            multi: true
        }
    ]
})
export class InputProductComponent implements OnInit, ControlValueAccessor {
    id = ""
    product: any = {}
    @Input() placeholder = ''
    private onChange!: any;

    constructor(private ms: NzModalService, private rs: SmartRequestService) {
    }

    ngOnInit(): void {
    }

    registerOnChange(fn: any): void {
        this.onChange = fn;
    }

    registerOnTouched(fn: any): void {
    }

    writeValue(obj: any): void {
        if (this.id !== obj) {
            this.id = obj
            if (this.id)
                this.load()
        }
    }

    load() {
        console.log('load product', this.id)
        this.rs.get('product/' + this.id).subscribe(res => {
            if (res.data) {
                this.product = res.data;
            }
        })
    }

    select() {
        this.ms.create({
            nzTitle: "选择",
            nzContent: ProductsComponent
        }).afterClose.subscribe(res => {
            console.log(res)
            if (res) {
                this.product = res
                this.id = res.id
                this.onChange(this.id)
            }
        })
    }

    change(value: string) {
        console.log('on change', value)
        this.id = value
        this.onChange(value)
        this.load()
    }
}
