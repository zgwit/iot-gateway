import { Component, Input, Output, EventEmitter } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
  selector: 'app-common-header',
  templateUrl: './common-header.component.html',
  styleUrls: ['./common-header.component.scss']
})
export class CommonHeaderComponent {
  downloadHref = "";
  importHref = "";
  @Input() title = "";
  @Input()
  set moduleName(moduleName: string) {
    this.downloadHref = `api/${moduleName}/export`;
    this.importHref = `api/${moduleName}/import`;
  };
  @Output() onLoad = new EventEmitter<string>();
  @Output() onSearch = new EventEmitter<string>();
  @Output() onAdd = new EventEmitter<string>();
  constructor(
    private msg: NzMessageService
  ) { }
  load() {
    this.onLoad.emit();
  }
  handleChange(info: any): void {
    if (info.file && info.file.response) {
      const res = info.file.response;
      if (!res.error) {
        this.msg.success(`成功导入${res.data}条数据!`);
        this.load();
      } else {
        this.msg.error(`${res.error}`);
      }
    }
  }
}
