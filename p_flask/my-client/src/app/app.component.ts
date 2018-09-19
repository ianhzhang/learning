import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {DomSanitizer,SafeResourceUrl} from '@angular/platform-browser'

// https://stackoverflow.com/questions/45781438/embed-pdf-in-angular2-component

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';
  docVisible = false;
  innerHtml

  constructor(private http: HttpClient, private sanitizer: DomSanitizer) {}

  ngOnInit() {
    var pdfurl = "http://localhost:5000/getPdf/2#view=FitH";

    this.innerHtml = this.sanitizer.bypassSecurityTrustHtml(
      "<object data='" + pdfurl + "' type='application/pdf' class='embed-responsive-item' width=100%; height=100%>" +
      "</object>");
  }


  load() {
    console.log("load");
    this.docVisible = true;
  }
}
