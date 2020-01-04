import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-master-list',
  templateUrl: './master-list.component.html',
  styleUrls: ['./master-list.component.css']
})
export class MasterListComponent implements OnInit {

  imageList: any = [];
  pageReady = false;
  updateTimerId = null;
  timeoutCnt = 0;

  constructor(private http: HttpClient) {
  }


  ngOnInit() {
    // period running this action
    this.updateTimerId = setInterval(() => {
      this.getInfoFromServer();
    }, 5000);
  }

  ngOnDestroy() {
    console.log("dashboard ngDestroy");
    if (this.updateTimerId) {
      clearInterval(this.updateTimerId)
    }
  }

  getInfoFromServer() {
    console.log(this.timeoutCnt);
    this.http.get("http://0.0.0.0:8085/imageInfoAll").subscribe(data => {
      this.imageList = data;
      console.log(this.imageList.length)
      if (this.imageList.length > 0) {
        this.pageReady = true;
        this.timeoutCnt = 0;
      } else {
        this.pageReady = false;
        this.timeoutCnt = this.timeoutCnt+1;
      }

      // from full name to find the filename. and put it in title field
      // remove "static/" from name field
      this.imageList.map(eachdata => {
        console.log(eachdata['name'])
        let nameItems = eachdata['name'].split("/");
        let lastIndex = nameItems.length - 1;

        eachdata['title'] = nameItems[nameItems.length - 1];

        // remove static/ prefix
        nameItems[0] = "";
        eachdata['name'] = nameItems.join('/');
        //eachdata['name'] = '/assets/' + eachdata['name'];
      });
    }, err => {
      this.pageReady = false;
      this.timeoutCnt = this.timeoutCnt+1;
    });

    //console.log(this.imageList)
  }


}
