unit excel2jsonminigame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Buttons, StdCtrls;

type

  { TExcel2JsonMiniGameFrame }

  TExcel2JsonMiniGameFrame = class(TFrame)
    GenerateButton: TButton;
    IndieFolderCheck: TCheckBox;
    ExcelDirLabel: TEdit;
    DistDirLabel: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    OpenExcelDirButton: TSpeedButton;
    OpenDistDirButton: TSpeedButton;
  private

  public

  end;

implementation

{$R *.lfm}

end.

