export class Version {
  version: string;
}

export class BackupResponse {
  log: string;
}

export class DiskStats {
  totalBytes: number;
  availableBytes: number;
}

export class Backup {
  name: string;
  sizeBytes: number;
  lastModified: number;
}

