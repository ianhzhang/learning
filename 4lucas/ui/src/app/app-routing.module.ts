import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { MasterListComponent } from './master-list/master-list.component';
import { DetailComponent } from './detail/detail.component';

const routes: Routes = [
  { path: '', component: MasterListComponent },
  { path: 'image/:index', component: DetailComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
