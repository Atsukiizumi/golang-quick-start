USE [master]
GO

/****** Object:  Database [gin_gorm_oj]    Script Date: 2024/8/12 11:37:35 ******/
CREATE DATABASE [gin_gorm_oj]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'gin_gorm_oj', FILENAME = N'[Your Real Path]\gin_gorm_oj.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 65536KB )
 LOG ON 
( NAME = N'gin_gorm_oj_log', FILENAME = N'[Your Real Path]\gin_gorm_oj_log.ldf' , SIZE = 8192KB , MAXSIZE = 2048GB , FILEGROWTH = 65536KB )
 WITH CATALOG_COLLATION = DATABASE_DEFAULT
GO

IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [gin_gorm_oj].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO

ALTER DATABASE [gin_gorm_oj] SET ANSI_NULL_DEFAULT OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET ANSI_NULLS OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET ANSI_PADDING OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET ANSI_WARNINGS OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET ARITHABORT OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET AUTO_CLOSE OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET AUTO_SHRINK OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET AUTO_UPDATE_STATISTICS ON 
GO

ALTER DATABASE [gin_gorm_oj] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET CURSOR_DEFAULT  GLOBAL 
GO

ALTER DATABASE [gin_gorm_oj] SET CONCAT_NULL_YIELDS_NULL OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET NUMERIC_ROUNDABORT OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET QUOTED_IDENTIFIER OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET RECURSIVE_TRIGGERS OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET  DISABLE_BROKER 
GO

ALTER DATABASE [gin_gorm_oj] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET TRUSTWORTHY OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET PARAMETERIZATION SIMPLE 
GO

ALTER DATABASE [gin_gorm_oj] SET READ_COMMITTED_SNAPSHOT OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET HONOR_BROKER_PRIORITY OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET RECOVERY FULL 
GO

ALTER DATABASE [gin_gorm_oj] SET  MULTI_USER 
GO

ALTER DATABASE [gin_gorm_oj] SET PAGE_VERIFY CHECKSUM  
GO

ALTER DATABASE [gin_gorm_oj] SET DB_CHAINING OFF 
GO

ALTER DATABASE [gin_gorm_oj] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO

ALTER DATABASE [gin_gorm_oj] SET TARGET_RECOVERY_TIME = 60 SECONDS 
GO

ALTER DATABASE [gin_gorm_oj] SET DELAYED_DURABILITY = DISABLED 
GO

ALTER DATABASE [gin_gorm_oj] SET ACCELERATED_DATABASE_RECOVERY = OFF  
GO

ALTER DATABASE [gin_gorm_oj] SET QUERY_STORE = OFF
GO

ALTER DATABASE [gin_gorm_oj] SET  READ_WRITE 
GO


USE [gin_gorm_oj]
GO
/****** Object:  Table [dbo].[category]    Script Date: 2024/8/12 11:36:52 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[category](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[identity] [varchar](36) NULL,
	[name] [varchar](100) NULL,
	[parent_id] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_at] [datetime] NULL,
	[deleted_at] [datetime] NULL,
 CONSTRAINT [PK_category] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[problem]    Script Date: 2024/8/12 11:36:52 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[problem](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[identity] [varchar](36) NULL,
	[category_id] [varchar](max) NULL,
	[title] [varchar](max) NULL,
	[content] [text] NULL,
	[max_men] [int] NULL,
	[max_runtime] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_at] [datetime] NULL,
	[deleted_at] [datetime] NULL,
 CONSTRAINT [PK_problem] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[submit]    Script Date: 2024/8/12 11:36:53 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[submit](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[identity] [varchar](36) NULL,
	[problem_identity] [varchar](36) NULL,
	[user_identity] [varchar](36) NULL,
	[status] [tinyint] NULL,
	[path] [text] NULL,
	[created_at] [datetime] NULL,
	[updated_at] [datetime] NULL,
	[deleted_at] [datetime] NULL,
 CONSTRAINT [PK_submit] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[user]    Script Date: 2024/8/12 11:36:53 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[user](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[identity] [varchar](36) NULL,
	[name] [varchar](100) NULL,
	[password] [varchar](32) NULL,
	[phone] [varchar](20) NULL,
	[mail] [varchar](100) NULL,
	[created_at] [datetime] NULL,
	[updated_at] [datetime] NULL,
	[deleted_at] [datetime] NULL,
 CONSTRAINT [PK_user] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[category] ADD  CONSTRAINT [DF_category_parent_id]  DEFAULT ((0)) FOR [parent_id]
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'唯一标识' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'category', @level2type=N'COLUMN',@level2name=N'identity'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'分类的名称' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'category', @level2type=N'COLUMN',@level2name=N'name'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'父级ID' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'category', @level2type=N'COLUMN',@level2name=N'parent_id'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'最大运行内存' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'problem', @level2type=N'COLUMN',@level2name=N'max_men'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'最大运行时长' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'problem', @level2type=N'COLUMN',@level2name=N'created_at'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'问题唯一标识' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'submit', @level2type=N'COLUMN',@level2name=N'problem_identity'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'用户唯一标识' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'submit', @level2type=N'COLUMN',@level2name=N'user_identity'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'提交状态' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'submit', @level2type=N'COLUMN',@level2name=N'status'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'代码路径' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'submit', @level2type=N'COLUMN',@level2name=N'path'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'用户名' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'user', @level2type=N'COLUMN',@level2name=N'name'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'密码' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'user', @level2type=N'COLUMN',@level2name=N'password'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'手机号' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'user', @level2type=N'COLUMN',@level2name=N'phone'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'邮箱' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'user', @level2type=N'COLUMN',@level2name=N'mail'
GO
