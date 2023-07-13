import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminComponent } from './admin/admin.component';
import { ProductComponent } from './admin/product/product.component';
import { ClientComponent } from './admin/client/client.component';
import { ClientEditComponent } from './admin/client/client-edit/client-edit.component';
import { ClientDetailComponent } from './admin/client/client-detail/client-detail.component';
import { SerialComponent } from './admin/serial/serial.component';
import { SerialDetailComponent } from './admin/serial/serial-detail/serial-detail.component';
import { SerialEditComponent } from './admin/serial/serial-edit/serial-edit.component';
import { LinkComponent } from './admin/link/link.component';
import { LinkEditComponent } from './admin/link/link-edit/link-edit.component';
import { LinkDetailComponent } from './admin/link/link-detail/link-detail.component';
import { ServerComponent } from './admin/server/server.component';
import { ServerEditComponent } from './admin/server/server-edit/server-edit.component';
import { ServerDetailComponent } from './admin/server/server-detail/server-detail.component';
import { DeviceComponent } from './admin/device/device.component';
import { DeviceEditComponent } from './admin/device/device-edit/device-edit.component';
import { ProductEditComponent } from './admin/product/product-edit/product-edit.component';

const pages: Routes = [
  {
    path: 'setting',
    loadChildren: () =>
      import('./admin/setting/setting.module').then((m) => m.SettingModule),
  },
];

const routes: Routes = [
  {
    path: 'admin',
    component: AdminComponent,
    children: [
      { path: 'product', component: ProductComponent },
      { path: 'product/edit/:id', component: ProductEditComponent },
      { path: 'create/product', component: ProductEditComponent },
      //{ path: 'product/:id', component: ProductDetailComponent },
      { path: 'device', component: DeviceComponent },
      { path: 'device/edit/:id', component: DeviceEditComponent },
      { path: 'create/device', component: DeviceEditComponent },
      //{ path: 'device/:id', component: DeviceDetailComponent },
      { path: 'server', component: ServerComponent },
      { path: 'server/edit/:id', component: ServerEditComponent },
      { path: 'create/server', component: ServerEditComponent },
      { path: 'server/:id', component: ServerDetailComponent },
      { path: 'link', component: LinkComponent },
      { path: 'link/:id', component: LinkDetailComponent },
      { path: 'link/edit/:id', component: LinkEditComponent },
      { path: 'create/link', component: LinkEditComponent },
      { path: 'serial', component: SerialComponent },
      { path: 'serial/:id', component: SerialDetailComponent },
      { path: 'serial/edit/:id', component: SerialEditComponent },
      { path: 'create/serial', component: SerialEditComponent },
      { path: 'client', component: ClientComponent },
      { path: 'client/edit/:id', component: ClientEditComponent },
      { path: 'create/client', component: ClientEditComponent },
      { path: 'client/:id', component: ClientDetailComponent },
      ...pages,
      { path: '**', redirectTo: 'device' },
    ],
  },
  {
    path: '**',
    redirectTo: 'admin',
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
