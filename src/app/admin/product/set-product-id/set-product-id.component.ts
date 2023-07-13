import { Component } from '@angular/core';
import { NzModalService } from 'ng-zorro-antd/modal';
import { ProductComponent } from '../product.component';
@Component({
  selector: 'app-set-product-id',
  templateUrl: './set-product-id.component.html',
  styleUrls: ['./set-product-id.component.scss']
})
export class SetProductIdComponent {
  product_id = '';
  constructor(
    private ms: NzModalService
  ) { }
  chooseProduct() {
    this.ms
      .create({
        nzTitle: '选择产品',
        nzWidth: '700px',
        nzContent: ProductComponent,
        nzFooter: null,
      })
      .afterClose.subscribe((product_id) => {
        this.product_id = product_id;
      });
  }
}
