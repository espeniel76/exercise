import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Body,
  Param,
  ParseIntPipe,
  HttpCode,
  HttpStatus,
} from '@nestjs/common';
import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { User } from './entities/user.entity';

@Controller('api/users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post()
  @HttpCode(HttpStatus.CREATED)
  async create(@Body() createUserDto: CreateUserDto): Promise<{ success: boolean; data: User }> {
    const user = await this.usersService.create(createUserDto);
    return { success: true, data: user };
  }

  @Get()
  async findAll(): Promise<{ success: boolean; data: User[] }> {
    const users = await this.usersService.findAll();
    return { success: true, data: users };
  }

  @Get(':id')
  async findOne(@Param('id', ParseIntPipe) id: number): Promise<{ success: boolean; data: User }> {
    const user = await this.usersService.findOne(id);
    return { success: true, data: user };
  }

  @Put(':id')
  async update(
    @Param('id', ParseIntPipe) id: number,
    @Body() updateUserDto: UpdateUserDto,
  ): Promise<{ success: boolean; data: User }> {
    const user = await this.usersService.update(id, updateUserDto);
    return { success: true, data: user };
  }

  @Delete(':id')
  async remove(@Param('id', ParseIntPipe) id: number): Promise<{ success: boolean; message: string }> {
    await this.usersService.remove(id);
    return { success: true, message: 'User deleted successfully' };
  }
}
