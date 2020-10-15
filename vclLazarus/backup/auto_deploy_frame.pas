unit auto_deploy_frame;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, ExtCtrls, ComCtrls, StdCtrls, Buttons;

type

  { TAutoDeploy }

  TAutoDeploy = class(TFrame)
    Button1: TButton;
    CloseButton1: TButton;
    ConfirmNameButton: TSpeedButton;
    ContextAdd: TButton;
    ContextPageControl: TPageControl;
    ContextPanel: TPanel;
    DeployContext: TTabSheet;
    DeployPanel: TPanel;
    FileContext: TListView;
    FileName: TEdit;
    GlobalContext: TListView;
    GlobalSheet: TTabSheet;
    ImageList: TImageList;
    KeyEdit: TEdit;
    KeyLabel: TLabel;
    KeyLabel1: TLabel;
    Label1: TLabel;
    LeftPanel: TPanel;
    BottomPanel: TPanel;
    ListBox1: TListBox;
    MainPanel: TPanel;
    NewButton: TSpeedButton;
    OpenDeployButton: TSpeedButton;
    Panel: TPanel;
    Panel1: TPanel;
    Panel2: TPanel;
    ProccedurePanel1: TPanel;
    ReverseNameButton: TSpeedButton;
    SaveButton: TSpeedButton;
    SaveButton1: TButton;
    SaveToButton: TSpeedButton;
    AddFileButton: TSpeedButton;
    RemoveButton: TSpeedButton;
    NewFolderButton: TSpeedButton;
    EditButton: TSpeedButton;
    Splitter1: TSplitter;
    Splitter2: TSplitter;
    Splitter3: TSplitter;
    Splitter4: TSplitter;
    TopPanel: TPanel;
    TopPanel1: TPanel;
    TreeView1: TTreeView;
    ValueEdit: TEdit;
    procedure LeftPanelClick(Sender: TObject);
    procedure MainPanelClick(Sender: TObject);
  private

  public

  end;

implementation

{$R *.lfm}

{ TAutoDeploy }

procedure TAutoDeploy.MainPanelClick(Sender: TObject);
begin

end;

procedure TAutoDeploy.LeftPanelClick(Sender: TObject);
begin

end;

end.

