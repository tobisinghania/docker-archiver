import { Injectable } from '@angular/core';
import {HttpClient, HttpEvent, HttpHeaders, HttpProgressEvent} from "@angular/common/http";
import {Backup, DiskStats, Version} from "./models";
import {Observable} from "rxjs";
import {catchError, filter, map} from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private httpClient: HttpClient) { }

  getVersion(): Observable<Version> {
    return this.httpClient.get<Version>('api/version');
  }


  getDiskStats(): Observable<DiskStats> {
    return this.httpClient.get<DiskStats>('api/diskStats');
  }

  getBackups(): Observable<Backup[]> {
    return this.httpClient.get<Backup[]>('api/backups');
  }

  createBackup(): Observable<void> {
    return this.httpClient.post<void>('api/backups', {});
  }

  downloadBackup(name: string) {
    window.open("static/backups/" + name, "blank")
  }

  deleteBackup(name: string): Observable<void> {
    return this.httpClient.delete<void>('api/backups/' + name)
  }

  restoreBackup(name: string): Observable<void> {
    return this.httpClient.post<void>('api/backups/restore/' + name, {})
  }

  postFile(fileToUpload: File): Observable<{loaded: number, total: number}> {
    const endpoint = 'api/uploadBackup';
    const formData: FormData = new FormData();
    formData.append('file', fileToUpload, fileToUpload.name);
    return this.httpClient
      .post(endpoint, formData, { headers: new HttpHeaders(), reportProgress: true, observe: "events" },)
      .pipe(
        filter(event => Object.getOwnPropertyNames(event).indexOf("loaded") >= 0),
        map((event: HttpProgressEvent) => {
          event = event as HttpProgressEvent;
          return {loaded: event.loaded, total: event.total}
        }),
        // catchError((e) => this.handleError(e));
      )
  }
}
