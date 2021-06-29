unit json_value_edit_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, StdCtrls;

type

  { TJsonValueEditFrame }

  TJsonValueEditFrame = class(TFrame)
    FormatCheck: TCheckBox;
    TypeList: TComboBox;
    KeyEdit: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Panel2: TPanel;
    Panel3: TPanel;
    ValueEdit: TMemo;
    procedure TypeListChange(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TJsonValueEditFrame }

procedure TJsonValueEditFrame.TypeListChange(Sender: TObject);
begin

end;

end.

