import { Module } from '@nestjs/common';
import { PrismaService } from '../prisma.service';
import { TelemetryController } from './telemetry.controller';
import { TelemetryRepository } from './telemetry.repository';
import { TelemetryService } from './telemetry.service';

@Module({
  controllers: [TelemetryController],
  providers: [TelemetryService, TelemetryRepository, PrismaService],
})
export class TelemetryModule {}
