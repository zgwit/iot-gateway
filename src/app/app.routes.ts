import {ActivatedRoute, Routes} from '@angular/router';
import {LoginComponent} from "./login/login.component";
import {UnknownComponent} from "@god-jason/smart";
import {authGuard} from "./auth.guard";
import {AdminComponent} from "./admin/admin.component";

export const routes: Routes = [
    {path: "", pathMatch: "full", redirectTo: "admin"},

    {path: "login", component: LoginComponent},
    {
        path: "admin",
        canActivate: [authGuard],
        component: AdminComponent,
        loadChildren: () => import('./admin/admin.module').then(m => m.AdminModule),
        data: {breadcrumb: "管理后台"}
    },

    {path: "**", component: UnknownComponent},
];

export function GetParentRouteUrl(route: ActivatedRoute): string {
    let base = route.snapshot.parent?.url.map(u => u.path).join("/") || 'admin'
    base = "/" + base
    console.log("base ", base)
    return base
}

export function GetParentRouteParam(route: ActivatedRoute, name: string) {
    return route.snapshot.parent?.paramMap.get(name)
}

export function GetParentRouteParamFilter(route: ActivatedRoute, name: string, key: string) {
    let filter: any = {}
    if (route.snapshot.parent?.paramMap.has(name)) {
        filter[key] = route.snapshot.parent?.paramMap.get(name)
    }
    return filter
}
