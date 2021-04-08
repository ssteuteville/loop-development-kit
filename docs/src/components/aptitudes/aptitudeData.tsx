import React from 'react';

export interface LDKLinks {
  node: string;
  go: string;
  dotnet: string;
}

export type LDKLinkActive = Partial<LDKLinks>;

export interface IAptitudeData {
  name: string;
  description: string;
  capabilities: ICapabilityData[];
  links?: LDKLinkActive;
}

export interface ICapabilityData {
  name: string;
  description: string;
  links?: LDKLinkActive;
}