import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavigationComponent } from './framework/navigation/navigation.component';
import { HeaderComponent } from './framework/header/header.component';
import { HomeComponent } from './features/home/home.component';
import { TaskModule } from './features/task/task.module';
import { QuestionnairesComponent } from './features/questionnaires/questionnaires.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    NavigationComponent,
    HomeComponent,
    QuestionnairesComponent,
  ],
  imports: [
    TaskModule,
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
