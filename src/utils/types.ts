export enum Environment {
  LOCAL = 'local',
  DEV = 'dev',
  STAGING = 'staging',
  PROD = 'prod',
}

export enum Service {
  DATABASE = 'database',
  STORAGE = 'storage',
}

export enum NotificationLevel {
  INFO = 'info',
  WARNING = 'warning',
  ERROR = 'error',
}

export interface Error {
  message: string;
  level: NotificationLevel;
}

export interface Metric {
  name: string;
  value: number;
  timestamp: Date;
}

export interface Deployment {
  id: string;
  service: Service;
  environment: Environment;
  startedAt: Date;
  finishedAt: Date;
}

export interface Release {
  id: string;
  service: Service;
  environment: Environment;
  version: string;
  releasedAt: Date;
}

export interface Cluster {
  name: string;
  environment: Environment;
  services: Service[];
}

export interface ServerlessConfig {
  functionName: string;
  region: string;
  envVariables: { [key: string]: string };
}