unit json_edit_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, ComCtrls, ValEdit, StdCtrls,
  Grids, PairSplitter, Buttons;

type

  { TJsonEditFrame }

  TJsonEditFrame = class(TFrame)
    ButtonImageList: TImageList;
    Button_Object: TSpeedButton;
    Button_Array: TSpeedButton;
    NewFileButton: TSpeedButton;
    OpenDirButton: TSpeedButton;
    Button_String: TSpeedButton;
    Button_Number: TSpeedButton;
    Button_Bool: TSpeedButton;
    Button_Null: TSpeedButton;
    UpButton: TSpeedButton;
    SaveToButton: TSpeedButton;
    SaveButton: TSpeedButton;
    Panel1: TPanel;
    EditPanel: TPanel;
    OpPanel: TPanel;
    ParsePanel: TPanel;
    Splitter1: TSplitter;
    Splitter2: TSplitter;
    StatusBar: TStatusBar;
    TreePanel: TTreeView;
    DownButton: TSpeedButton;
    procedure Button_ObjectClick(Sender: TObject);
    procedure EditPanelClick(Sender: TObject);
    procedure PairSplitterSide2MouseDown(Sender: TObject; Button: TMouseButton;
      Shift: TShiftState; X, Y: Integer);
  private

  public

  end;

implementation

{$R *.lfm}

{ TJsonEditFrame }

procedure TJsonEditFrame.EditPanelClick(Sender: TObject);
begin

end;

procedure TJsonEditFrame.Button_ObjectClick(Sender: TObject);
begin

end;

procedure TJsonEditFrame.PairSplitterSide2MouseDown(Sender: TObject;
  Button: TMouseButton; Shift: TShiftState; X, Y: Integer);
begin

end;

end.

