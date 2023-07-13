import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Router } from '@angular/router';
import { NzMessageService } from 'ng-zorro-antd/message';
import { RequestService } from 'src/app/request.service';

@Component({
    selector: 'app-operate',
    templateUrl: './operate.component.html',
    styleUrls: ['./operate.component.scss'],
})
export class OperateComponent {
  constructor(private rs: RequestService, private msg: NzMessageService, private router: Router,){}
    @Input() id: any;
    @Input() url: any;
    @Output ()onSend = new EventEmitter<any>();
     
    edit() {
      const path = `/admin/${this.url}/edit/${this.id}`;
      this.router.navigateByUrl(path);
    }

    delete() {
      this.rs.get(`${this.url}/${this.id}/delete`).subscribe((res) => {
        this.msg.success('删除成功'); 
        this.onSend.emit();
      });

     
    }

    disable() {
      this.rs.get(`${this.url}/${this.id}/disable`).subscribe((res) => {
        this.msg.success(`已禁用!`);
        this.onSend.emit();
       
      });
    }
    enable() {
      this.rs.get(`${this.url}/${this.id}/enable`).subscribe((res) => {
        this.msg.success(`已启用!`);
        this.onSend.emit();
     
      });
 
    }
    start() {
      this.rs.get(`${this.url}/${this.id}/start`).subscribe((res) => {
        this.msg.success(`已启动!`);
        this.onSend.emit();
      
      });
    }
    stop() {
      this.rs.get(`${this.url}/${this.id}/stop`).subscribe((res) => {
        this.msg.success(`已停止!`);
        this.onSend.emit();
       
      });
    }
}
