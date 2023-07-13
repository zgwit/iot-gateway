import { Component, Input, ViewChild } from '@angular/core';
import { RequestService } from '../../request.service';
import { Router, NavigationExtras } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
  selector: 'app-tunnel-device',
  templateUrl: './tunnel-device.component.html',
  styleUrls: ['./tunnel-device.component.scss'],
})
export class TunnelDeviceComponent {
  _tunnel = '';

  @Input()
  set tunnel(id: any) {
    this._tunnel = id;
    this.query.filter = { tunnel_id: this._tunnel };
    this.load();
  }

  constructor(
    private router: Router,
    private rs: RequestService,
    private msg: NzMessageService
  ) {
    //this.load();
  }


  isVisible!: boolean;
  addVisible = false;
  loading = true;
  datum: any[] = [];
  total = 1;
  pageSize = 20;
  pageIndex = 1;
  query: any = {};

  clientFm(num: number) {
    if (num) this.load();
    this.isVisible = false;
  }

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
    this.datum.splice(index, 1);
    this.rs.get(`device/${id}/delete`).subscribe((res) => {
      this.msg.success('删除成功');
      this.isVisible = false;
      this.load();
    });
  }

  add() {
    const navigationExtras: NavigationExtras = {
      queryParams: { 'tunnelId': this._tunnel }
    };
    this.router.navigate([`/admin/create/device`], navigationExtras);
  }

  edit(id: number, data: any) {
    const path = `/admin/device/edit/${id}`;
    this.router.navigateByUrl(path);
  }

  search(text: any) {
    if (text)
      this.query.filter = {
        tunnel_id: this._tunnel,
        id: text,
      };
    else this.query.filter = { tunnel_id: this._tunnel };
    this.load();
  }

  cancel() {
    this.msg.info('取消删除');
  }

  open(id: string) {
    this.router.navigateByUrl('/admin/device/' + id);
  }
}
