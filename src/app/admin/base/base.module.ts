import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from "@angular/forms";
import { ReactiveFormsModule } from "@angular/forms";
import { DragDropModule } from '@angular/cdk/drag-drop'

import { NzTableModule } from "ng-zorro-antd/table";
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzFormModule } from "ng-zorro-antd/form";
import { NzSelectModule } from "ng-zorro-antd/select";
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzIconModule } from "ng-zorro-antd/icon";
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzSwitchModule } from 'ng-zorro-antd/switch';
import { NzPageHeaderModule } from 'ng-zorro-antd/page-header';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzUploadModule } from 'ng-zorro-antd/upload';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { SearchBoxComponent } from './search-box/search-box.component';
import { EditTableComponent } from './edit-table/edit-table.component';
import { CommonHeaderComponent } from './common-header/common-header.component';
import { OperateComponent } from './operate/operate.component';
import { NzDividerModule } from 'ng-zorro-antd/divider';

@NgModule({
  declarations: [
    SearchBoxComponent,
    EditTableComponent,
    CommonHeaderComponent,
    OperateComponent,
  ],
  exports: [
    SearchBoxComponent,
    EditTableComponent,
    CommonHeaderComponent,
    OperateComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    NzInputModule,
    NzDropDownModule ,
    NzButtonModule,
    NzDividerModule,
    NzTableModule,
    NzFormModule,
    ReactiveFormsModule,
    DragDropModule,
    NzSelectModule,
    NzSpaceModule,
    NzIconModule,
    NzInputNumberModule,
    NzSwitchModule,
    NzPageHeaderModule,
    NzLayoutModule,
    NzUploadModule,
  ]
})
export class BaseModule { }
