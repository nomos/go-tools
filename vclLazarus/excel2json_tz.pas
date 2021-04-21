unit excel2json_tz;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Buttons, StdCtrls;

type

  { TExcel2JsonTz }

  TExcel2JsonTz = class(TFrame)
    ExportSelect: TComboBox;
    DistDirLabel: TEdit;
    ExcelDirLabel: TEdit;
    GenerateButton: TButton;
    HelpButton: TSpeedButton;
    Label1: TLabel;
    Label2: TLabel;
    OpenDistDirButton: TSpeedButton;
    OpenExcelDirButton: TSpeedButton;
    StaticText1: TStaticText;
    procedure ExportSelectChange(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TExcel2JsonTz }

procedure TExcel2JsonTz.ExportSelectChange(Sender: TObject);
begin

end;

end.

