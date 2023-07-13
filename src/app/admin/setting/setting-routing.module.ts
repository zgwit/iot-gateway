import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { WebComponent } from "./web/web.component";
import { DatabaseComponent } from "./database/database.component";
import { LogComponent } from "./log/log.component";
import { MqttComponent } from './mqtt/mqtt.component';
// import { PageNotFoundComponent } from "../base/page-not-found/page-not-found.component";
const routes: Routes = [
    { path: '', pathMatch: "full", redirectTo: "web" },
    { path: 'web', component: WebComponent },
    { path: 'database', component: DatabaseComponent },
    { path: 'log', component: LogComponent },
    { path: 'mqtt', component: MqttComponent },
    // { path: '**', component: PageNotFoundComponent }
];
@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class SettingRoutingModule {
}