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
    // 0.0.0.0:8080 is external ip
    // http://0.0.0.0:8080/api/hello works
    // http://0.0.0.0:5000/api/hello should not work. browser is outside container and cannot access to container
    // internal makhttp://python_my-server_1:5000/api/hello not work.
    this.http.get("http://0.0.0.0:5001/api/hello").subscribe(data=> {
      console.log(data)
      this.result = data["rslt"]
    })
  }
}
