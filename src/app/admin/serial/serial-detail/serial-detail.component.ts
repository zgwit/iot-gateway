import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { RequestService } from "../../../request.service";

@Component({
  selector: 'app-serial-detail',
  templateUrl: './serial-detail.component.html',
  styleUrls: ['./serial-detail.component.scss']
})
export class SerialDetailComponent implements OnInit {
  id: string = ""
  data: any = {}

  constructor(private route: ActivatedRoute, private rs: RequestService) { }

  ngOnInit(): void {
    // @ts-ignore
    this.id = this.route.snapshot.paramMap.get("id")

    this.load()
  }

  load() {
    this.rs.get(`serial/${this.id}`).subscribe(res => {
      this.data = res.data;
    })
  }
}
