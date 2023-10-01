USE [BookingServices]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[TypeBusiness](
	[TypeBusinessID] [bigint] IDENTITY(1,1) NOT NULL,
	[TypeBusinessName ] [nvarchar](100) NULL,
	[Description ] [nvarchar](100) NOT NULL,
	[NameServiceProducers] [nvarchar](100) NOT NULL,
	[UseMultipleSlotBooking] [bit] NOT NULL,
	[MarkDeletion] [bit] NULL,
	[UseSelectSlotService] [bit] NOT NULL,
 CONSTRAINT [PK_TypeBusiness] PRIMARY KEY CLUSTERED 
(
	[TypeBusinessID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]

INSERT INTO [dbo].[TypeBusiness]
           ([TypeBusinessName ]
           ,[Description ]
           ,[NameServiceProducers]
           ,[UseMultipleSlotBooking]
           ,[MarkDeletion]
           ,[UseSelectSlotService])
     VALUES
           ('Парикмахерская'
           ,'Здесь стригут людей'
           ,'Парикмахеры'
           ,0
           ,0
           ,0),
            ('Автомойка'
           ,'Здесь моют автомобили'
           ,'Боксы'
           ,0
           ,0
           ,0)

CREATE TABLE [dbo].[UsersRoles](
	[RoleID] [int] IDENTITY(1,1) NOT NULL,
	[name] [nvarchar](50) NULL,
	[description] [nvarchar](250) NULL,
 CONSTRAINT [PK_UsersRoles] PRIMARY KEY CLUSTERED 
(
	[RoleID] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]

INSERT INTO [dbo].[UsersRoles]
           ([name]
           ,[description])
     VALUES
           ('Admin'
           ,'Администратор приложения'),

           ('ServiceProvider'
           ,'Поставщик услуг'),

           ('ServiceСonsumer'
           ,'Потребитель услуг')


CREATE TABLE Users (
  UserID INT IDENTITY(1,1) PRIMARY KEY CLUSTERED,
  UserName NVARCHAR(50),
  Password NVARCHAR(50),
  RoleID INT,
  Email NVARCHAR(256),
  Phone NVARCHAR(20),
  INDEX IX_RoleID (RoleID),
  FOREIGN KEY (RoleID) REFERENCES UsersRoles(RoleID)
);

INSERT INTO [dbo].[Users]
           ([UserName]
           ,Password
           ,RoleID
           ,Email
           ,Phone
           )
     VALUES
           ('Admin'
           ,''
           ,1
           ,'admin@bookingservice.com'
           ,'89169370850'),

      ('Service provider'
           ,''
           ,1
           ,'admin@company.com'
           ,'89169370851'),

    ('Service consumer'
           ,''
           ,1
           ,'ivan@company.com'
           ,'89169370852')