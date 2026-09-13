import { Injectable } from '@nestjs/common';
import { randomUUID } from 'crypto';
import { Telemetry } from './telemetry.entity';
import { TelemetryRepository } from './telemetry.repository';

@Injectable()
export class TelemetryService {
  constructor(private readonly telemetryRepository: TelemetryRepository) {}

  async create(
    deviceId: string,
    type: string,
    value: string,
  ): Promise<Telemetry> {
    const telemetry: Telemetry = {
      id: randomUUID(),
      deviceId,
      type,
      value,
      timestamp: new Date(),
    };

    return this.telemetryRepository.save(telemetry);
  }

  async findAll(): Promise<Telemetry[]> {
    return this.telemetryRepository.findAll();
  }
}
