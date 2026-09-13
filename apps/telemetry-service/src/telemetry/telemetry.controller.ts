import { Body, Controller, Get, Post } from '@nestjs/common';

import { TelemetryService } from './telemetry.service';

@Controller('api/v1/telemetry')
export class TelemetryController {
  constructor(private readonly telemetryService: TelemetryService) {}

  @Post()
  create(
    @Body()
    body: {
      deviceId: string;
      type: string;
      value: string;
    },
  ) {
    return this.telemetryService.create(body.deviceId, body.type, body.value);
  }

  @Get()
  findAll() {
    return this.telemetryService.findAll();
  }
}
