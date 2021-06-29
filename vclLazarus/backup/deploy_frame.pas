unit deploy_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, StdCtrls, ComCtrls, Buttons;

type

  { TDeployFrame }

  TDeployFrame = class(TFrame)
    Button1: TButton;
    CloseButton1: TButton;
    ConfirmNameButton: TSpeedButton;
    ContextAdd: TButton;
    ContextPageControl: TPageControl;
    ContextPanel: TPanel;
    DeployContext: TTabSheet;
    DeployPanel: TPanel;
    FileContextList: TListView;
    FileName: TEdit;
    GlobalContextList: TListView;
    GlobalSheet: TTabSheet;
    KeyEdit: TEdit;
    KeyLabel: TLabel;
    KeyLabel1: TLabel;
    Label1: TLabel;
    ListBox1: TListBox;
    OpenValueFolderButton: TSpeedButton;
    Panel1: TPanel;
    Panel2: TPanel;
    Panel3: TPanel;
    Panel4: TPanel;
    Panel5: TPanel;
    ProccedurePanel1: TPanel;
    ReverseNameButton: TSpeedButton;
    SaveButton1: TButton;
    Splitter2: TSplitter;
    Splitter4: TSplitter;
    TopPanel1: TPanel;
    ValueEdit: TEdit;
  private

  public

  end;

implementation

{$R *.lfm}

end.

