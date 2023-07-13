import { RequestService } from '../../request.service';
import { Component, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from '../base/table';
@Component({
  selector: 'app-device',
  templateUrl: './device.component.html',
  styleUrls: ['./device.component.scss'],
})
export class DeviceComponent {
  constructor(
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    this.load();
  }

  addVisible = false;
  loading = true;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  load() {
    this.loading = true;
    this.rs
      .post('device/search', this.query)
      .subscribe((res) => {
        this.datum = res.data;
        this.total = res.total;
      })
      .add(() => {
        this.loading = false;
      });
  }
  delete(index: number, id: number) {
    this.rs.get(`device/${id}/delete`).subscribe((res) => {
      this.msg.success('删除成功');
      this.load();
    });
  }
  add() {
    this.router.navigateByUrl(`/admin/create/device`);
  }
  edit(id: number, data: any) {
    const path = `/admin/device/edit/${id}`;
    this.router.navigateByUrl(path);
  }
  onQuery($event: NzTableQueryParams) {
    ParseTableQuery($event, this.query);
    this.load();
  }
  pageIndexChange(pageIndex: number) {
    this.query.skip = pageIndex - 1;
  }
  pageSizeChange(pageSize: number) {
    this.query.limit = pageSize;
  }
  search(text: any) {
    if (text)
      this.query.filter = {
        id: text,
      };
    else this.query = {};
    this.load();
  }
  cancel() {
    this.msg.info('取消删除');
  }

  open(id: string) {
    //this.router.navigateByUrl("/admin/device/" + id)
  }
  handleToggleStatus(index: number, data: { disabled: boolean, id: number }) {
    const { disabled, id } = data;
    const url = disabled ? `device/${id}/enable` : `device/${id}/disable`;
    this.rs.get(url).subscribe((res) => {
      this.msg.success(`${disabled ? '启用' : '禁用'}成功!`);
      this.load();
    });
  }
}
