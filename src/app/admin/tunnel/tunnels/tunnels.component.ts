import {Component} from '@angular/core';
import {NzTabsModule} from "ng-zorro-antd/tabs";
import {LinksComponent} from "../../link/links/links.component";
import {ClientsComponent} from "../../client/clients/clients.component";
import {SerialsComponent} from "../../serial/serials/serials.component";

@Component({
    selector: 'app-tunnels',
    standalone: true,
    imports: [
        NzTabsModule,
        LinksComponent,
        ClientsComponent,
        SerialsComponent
    ],
    templateUrl: './tunnels.component.html',
    styleUrl: './tunnels.component.scss'
})
export class TunnelsComponent {

}
