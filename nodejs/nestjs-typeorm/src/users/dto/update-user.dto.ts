import { IsEmail, IsOptional, IsString, MinLength, MaxLength } from 'class-validator';

export class UpdateUserDto {
  @IsEmail()
  @IsOptional()
  email?: string;

  @IsString()
  @IsOptional()
  @MinLength(3)
  @MaxLength(50)
  username?: string;

  @IsString()
  @IsOptional()
  @MinLength(6)
  password?: string;
}
