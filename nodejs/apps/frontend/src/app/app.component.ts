import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { map, take } from 'rxjs/operators';

@Component({
  selector: 'nodejs-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  ocean = '';

  constructor(private readonly http: HttpClient) {
  }

  ngOnInit(): void {
    this.http.get<{ message: string }>('http://localhost:8003/api')
      .pipe(take(1), map(cur => cur.message))
      .subscribe(res => this.ocean = res);
  }
}
