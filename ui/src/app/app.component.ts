import {Component} from '@angular/core';
import {ApiService} from "../core/api/api.service";
import {BehaviorSubject, interval, Observable, of, throwError} from "rxjs";
import {catchError, finalize, map, shareReplay, startWith, switchMap, tap} from "rxjs/operators";
import {Backup, DiskStats} from "../core/api/models";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'backup-ui';

  version$: Observable<string>;
  diskStats$: Observable<DiskStats>;
  diskStatsPercentage$: Observable<string>;
  backups$: Observable<Backup[]>;

  createBackupButtonDisabled$ = new BehaviorSubject(false);
  createBackupStatusMessage$ = new BehaviorSubject(null);
  createBackupState$ = new BehaviorSubject<'pending' | 'failed' | 'success'>(null);

  restoreBackupButtonDisabled$ = new BehaviorSubject(false);
  restoreBackupStatusMessage$ = new BehaviorSubject(null);
  restoreBackupState$ = new BehaviorSubject<'pending' | 'failed' | 'success'>(null);


  fileToUpload: File = null;
  lastRestored: string;

  constructor(private api: ApiService, private snackBar: MatSnackBar) {
    this.version$ = this.api.getVersion().pipe(
      tap(it => console.log("result", it, it.version)),
      map(it => it.version)
    );

    this.diskStats$ = interval(1000).pipe(startWith(0))
      .pipe(
        switchMap(() => this.api.getDiskStats().pipe(catchError(() => of(null)))),
        shareReplay(1)
      );

    this.diskStatsPercentage$ = this.diskStats$.pipe(
      map(it =>
        Math.round((it.availableBytes / it.totalBytes * 100 + Number.EPSILON) * 100) / 100 + '%')
    )

    this.backups$ = interval(1000).pipe(startWith(0))
      .pipe(
        switchMap(() => this.api.getBackups().pipe(catchError(() => of([])))),

        shareReplay(1)
      )

  }

  handleFileInput(files: FileList) {
    this.fileToUpload = files.item(0);
  }

  uploadBackup() {
    this.api.postFile(this.fileToUpload).subscribe(data => {
      // do something, if upload success
    }, error => {
      console.log(error);
    });
  }

  createNewBackup() {
    this.createBackupStatusMessage$.next(false);
    this.createBackupButtonDisabled$.next(true);
    this.createBackupState$.next('pending')
    this.api.createBackup()
      .pipe(
        catchError(e => {
          this.createBackupStatusMessage$.next('Backupcreation failed: ' + e);
          this.createBackupState$.next('failed')
          return throwError(e);
        }),
        finalize(() => this.createBackupButtonDisabled$.next(false)),
      ).subscribe(() => {
      this.createBackupStatusMessage$.next('Successfully created backup');
      this.createBackupState$.next('success')
    }, error => {
    })
  }

  restoreBackup(name: string) {
    this.lastRestored = name;

    this.restoreBackupStatusMessage$.next("Restoring...");
    this.restoreBackupButtonDisabled$.next(true);
    this.restoreBackupState$.next('pending')
    this.api.restoreBackup(name)
      .pipe(
        catchError(e => {
          this.restoreBackupStatusMessage$.next('Restoring backup failed: ' + e.statusText);
          this.restoreBackupState$.next('failed')
          return throwError(e);
        }),
        finalize(() => this.restoreBackupButtonDisabled$.next(false)),
      ).subscribe(() => {
      this.restoreBackupStatusMessage$.next('Successfully restored backup');
      this.restoreBackupState$.next('success')
    }, error => {
    })
  }

  downloadBackup(name: string) {
    this.api.downloadBackup(name);
  }

  deleteBackup(name: string) {
    this.api.deleteBackup(name)
      .subscribe(value => {
        this.snackBar.open("Successfully delete backup " + name)
      }, error => this.snackBar.open("Error deleting backup " + name));
  }


}
