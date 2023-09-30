USE [BookingServices]
GO

USE [BookingServices]
GO

/****** Object:  Table [dbo].[TypeBusiness]    Script Date: 29.09.2023 18:27:09 ******/
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
GO

;
USE [BookingServices]
GO

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
GO


