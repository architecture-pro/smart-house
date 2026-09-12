import { Injectable } from '@nestjs/common';
import { PrismaService } from '../prisma.service';
import { Telemetry } from './telemetry.entity';

@Injectable()
export class TelemetryRepository {
  constructor(private readonly prisma: PrismaService) {}

  async save(telemetry: Telemetry): Promise<Telemetry> {
    return this.prisma.telemetry.create({
      data: {
        id: telemetry.id,
        deviceId: telemetry.deviceId,
        type: telemetry.type,
        value: telemetry.value,
        timestamp: telemetry.timestamp,
      },
    });
  }

  async findAll(): Promise<Telemetry[]> {
    return this.prisma.telemetry.findMany();
  }
}
