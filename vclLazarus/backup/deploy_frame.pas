unit deploy_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, ComCtrls, StdCtrls, Buttons;

type

  { TDeployFrame }

  TDeployFrame = class(TFrame)
    Button1: TButton;
    ContextAdd: TButton;
    KeyEdit: TEdit;
    ValueEdit: TEdit;
    FileContext: TListView;
    KeyLabel1: TLabel;
    Label1: TLabel;
    KeyLabel: TLabel;
    Panel1: TPanel;
    Panel2: TPanel;
    SaveButton: TButton;
    CloseButton1: TButton;
    ContextPageControl: TPageControl;
    ContextPanel: TPanel;
    DeployContext: TTabSheet;
    DeployPanel: TPanel;
    FileName: TEdit;
    GlobalSheet: TTabSheet;
    ListBox1: TListBox;
    GlobalContext: TListView;
    ConfirmNameButton: TSpeedButton;
    ReverseNameButton: TSpeedButton;
    Splitter1: TSplitter;
    TopPanel: TPanel;
    ProccedurePanel1: TPanel;
    Splitter2: TSplitter;
    procedure GlobalContextAddClick(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TDeployFrame }

procedure TDeployFrame.GlobalContextAddClick(Sender: TObject);
begin

end;

end.

