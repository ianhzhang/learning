import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {DomSanitizer,SafeResourceUrl} from '@angular/platform-browser'



@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';
  result="notYet"


  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.result="notYet"

  }


  load() {
    console.log("load");
    this.http.get("http://0.0.0.0:5001/hello").subscribe(data=> {
      console.log(data)
      this.result = data["rslt"]
    })
  }
}
