import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, RouterState} from "@angular/router";
import {RequestService} from "../../../request.service";

@Component({
  selector: 'app-link-detail',
  templateUrl: './link-detail.component.html',
  styleUrls: ['./link-detail.component.scss']
})
export class LinkDetailComponent implements OnInit {
  id: string = ""
  data: any = {}

  constructor(private route: ActivatedRoute, private rs: RequestService) {
  }

  ngOnInit(): void {
    // @ts-ignore
    this.id = this.route.snapshot.paramMap.get("id")

    this.load()
  }

  load() {
    this.rs.get(`link/${this.id}`).subscribe(res => {
      this.data = res.data;
    })
  }

}
