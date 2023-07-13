import { RequestService } from './../../request.service';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzTableQueryParams } from 'ng-zorro-antd/table';
import { ParseTableQuery } from '../base/table';
@Component({
  selector: 'app-serial',
  templateUrl: './serial.component.html',
  styleUrls: ['./serial.component.scss'],
})
export class SerialComponent {
  constructor(
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    this.load();
  }

  loading = true;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};
  load() {
    this.loading = true;
    this.rs
      .post('serial/search', this.query)
      .subscribe((res) => {
        this.datum = res.data;
        this.total = res.total;
      })
      .add(() => {
        this.loading = false;
      });
  }

  delete(index: number, id: number) {
    this.datum.splice(index, 1);
    this.rs.get(`serial/${id}/delete`).subscribe((res) => {
      this.msg.success('删除成功');
      this.load();
    });
  }
  status(num: number, id: any) {
    if (num) {
      this.rs.get(`serial/${id}/start`).subscribe((res) => {
        this.msg.success(`已启动!`);
        this.load();
      });
    }
    else {
      this.rs.get(`serial/${id}/stop`).subscribe((res) => {
        this.msg.success(`已停止!`);
        this.load();
      });
    }
  }
  add() {
    this.router.navigateByUrl(`/admin/create/serial`);
  }
  edit(id: number, data: any) {
    const path = `/admin/serial/edit/${id}`;
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
    this.router.navigateByUrl('/admin/serial/' + id);
  }
  handleToggleStatus(index: number, data: { disabled: boolean, id: number }) {
    const { disabled, id } = data;
    const url = disabled ? `serial/${id}/enable` : `serial/${id}/disable`;
    this.rs.get(url).subscribe((res) => {
      this.msg.success(`${disabled ? '启用' : '禁用'}成功!`);
      this.load();
    });
  }
}
