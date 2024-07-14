import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {DashComponent} from "./dash/dash.component";
import {ProductsComponent} from "./product/products/products.component";
import {ProductEditComponent} from "./product/product-edit/product-edit.component";
import {ProductDetailComponent} from "./product/product-detail/product-detail.component";
import {DevicesComponent} from "./device/devices/devices.component";
import {DeviceEditComponent} from "./device/device-edit/device-edit.component";
import {DeviceDetailComponent} from "./device/device-detail/device-detail.component";
import {ServersComponent} from "./server/servers/servers.component";
import {ServerEditComponent} from "./server/server-edit/server-edit.component";
import {ServerDetailComponent} from "./server/server-detail/server-detail.component";
import {LinksComponent} from "./link/links/links.component";
import {LinkEditComponent} from "./link/link-edit/link-edit.component";
import {LinkDetailComponent} from "./link/link-detail/link-detail.component";
import {SerialsComponent} from "./serial/serials/serials.component";
import {SerialEditComponent} from "./serial/serial-edit/serial-edit.component";
import {SerialDetailComponent} from "./serial/serial-detail/serial-detail.component";
import {ClientsComponent} from "./client/clients/clients.component";
import {ClientEditComponent} from "./client/client-edit/client-edit.component";
import {ClientDetailComponent} from "./client/client-detail/client-detail.component";
import {SettingComponent} from "./setting/setting.component";
import {UnknownComponent} from "@god-jason/smart";

const routes: Routes = [
    {path: "", pathMatch: "full", redirectTo: "dash"},
    {path: 'dash', component: DashComponent, title: "控制台", data: {breadcrumb: "控制台"}},

    {path: 'product', component: ProductsComponent, title: "产品列表", data: {breadcrumb: "产品列表"}},
    {path: 'product/create', component: ProductEditComponent, title: "创建产品", data: {breadcrumb: "创建产品"}},
    {path: 'product/:id', component: ProductDetailComponent, title: "产品详情", data: {breadcrumb: "产品详情"}},
    {path: 'product/:id/edit', component: ProductEditComponent, title: "编辑产品", data: {breadcrumb: "编辑产品"}},

    {path: 'device', component: DevicesComponent, title: "设备列表", data: {breadcrumb: "设备列表"}},
    {path: 'device/create', component: DeviceEditComponent, title: "创建设备", data: {breadcrumb: "创建设备"}},
    {path: 'device/:id', component: DeviceDetailComponent, title: "设备详情", data: {breadcrumb: "设备详情"}},
    {path: 'device/:id/edit', component: DeviceEditComponent, title: "编辑设备", data: {breadcrumb: "编辑设备"}},

    {path: 'server', component: ServersComponent, title: "TCP服务器列表", data: {breadcrumb: "TCP服务器列表"}},
    {
        path: 'server/create',
        component: ServerEditComponent,
        title: "创建TCP服务器",
        data: {breadcrumb: "创建TCP服务器"}
    },
    {path: 'server/:id', component: ServerDetailComponent, title: "TCP服务器详情", data: {breadcrumb: "TCP服务器详情"}},
    {
        path: 'server/:id/edit',
        component: ServerEditComponent,
        title: "TCP服务器编辑",
        data: {breadcrumb: "TCP服务器编辑"}
    },

    {path: 'link', component: LinksComponent, title: "TCP连接列表", data: {breadcrumb: "TCP连接列表"}},
    {path: 'link/create', component: LinkEditComponent, title: "创建TCP连接", data: {breadcrumb: "创建TCP连接"}},
    {path: 'link/:id', component: LinkDetailComponent, title: "TCP连接详情", data: {breadcrumb: "TCP连接详情"}},
    {path: 'link/:id/edit', component: LinkEditComponent, title: "TCP连接编辑", data: {breadcrumb: "TCP连接编辑"}},

    {path: 'serial', component: SerialsComponent, title: "串口列表", data: {breadcrumb: "串口列表"}},
    {path: 'serial/create', component: SerialEditComponent, title: "创建串口", data: {breadcrumb: "创建串口"}},
    {path: 'serial/:id', component: SerialDetailComponent, title: "串口详情", data: {breadcrumb: "串口详情"}},
    {path: 'serial/:id/edit', component: SerialEditComponent, title: "串口编辑", data: {breadcrumb: "串口编辑"}},

    {path: 'client', component: ClientsComponent, title: "客户端列表", data: {breadcrumb: "客户端列表"}},
    {path: 'client/create', component: ClientEditComponent, title: "创建客户端", data: {breadcrumb: "创建客户端"}},
    {path: 'client/:id', component: ClientDetailComponent, title: "客户端详情", data: {breadcrumb: "客户端详情"}},
    {path: 'client/:id/edit', component: ClientEditComponent, title: "客户端编辑", data: {breadcrumb: "客户端编辑"}},

    {path: "setting", component: SettingComponent, title: "设置", data: {breadcrumb: "设置"}},

    {path: "**", component: UnknownComponent},
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class AdminRoutingModule {
}
