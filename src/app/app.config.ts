import {ApplicationConfig, importProvidersFrom, LOCALE_ID} from '@angular/core';
import {provideRouter} from '@angular/router';

import {routes} from './app.routes';
import {provideNzI18n, zh_CN} from 'ng-zorro-antd/i18n';
import {registerLocaleData} from '@angular/common';
import {FormsModule} from '@angular/forms';
import zh from '@angular/common/locales/zh';
import {provideAnimations} from '@angular/platform-browser/animations';

import {NZ_ICONS} from "ng-zorro-antd/icon";
import {IconDefinition} from '@ant-design/icons-angular';
import {
    ApartmentOutline,
    AppstoreAddOutline,
    AppstoreOutline,
    ArrowLeftOutline,
    BackwardOutline,
    BellOutline,
    BlockOutline,
    BuildOutline,
    CloseCircleOutline,
    ClusterOutline,
    ControlOutline,
    DashboardOutline,
    DeleteOutline,
    DisconnectOutline,
    DownloadOutline,
    DragOutline,
    EditOutline,
    ExportOutline,
    EyeOutline,
    ImportOutline,
    LinkOutline,
    LockOutline,
    MenuFoldOutline,
    MenuUnfoldOutline,
    PlayCircleOutline,
    PlusOutline,
    ProfileOutline,
    ReloadOutline,
    SettingOutline,
    UploadOutline,
    UserOutline,
    VideoCameraOutline,
} from '@ant-design/icons-angular/icons';
import {provideEcharts} from "ngx-echarts";
import {HttpClientModule, provideHttpClient} from "@angular/common/http";

registerLocaleData(zh);

const icons: IconDefinition[] = [
    MenuFoldOutline,
    MenuUnfoldOutline,
    DashboardOutline,
    PlusOutline,
    BellOutline,
    SettingOutline,
    EditOutline,
    ApartmentOutline,
    BlockOutline,
    AppstoreOutline,
    AppstoreAddOutline,
    DeleteOutline,
    DownloadOutline,
    UploadOutline,
    UserOutline,
    ProfileOutline,
    EyeOutline,
    ReloadOutline,
    BackwardOutline,
    ArrowLeftOutline,
    LockOutline,
    DisconnectOutline,
    LinkOutline,
    DragOutline,
    ExportOutline,
    ImportOutline,
    VideoCameraOutline,
    ClusterOutline,
    PlayCircleOutline,
    CloseCircleOutline,
    BuildOutline,
    ControlOutline,
];

export const appConfig: ApplicationConfig = {
    providers: [
        provideRouter(routes),
        provideNzI18n(zh_CN),
        provideHttpClient(),
        provideAnimations(),
        provideEcharts(),
        {provide: NZ_ICONS, useValue: icons},
        {provide: LOCALE_ID, useValue: "zh_CN"},
        //{provide: API_BASE, useValue: "/api/"},
    ]
};
